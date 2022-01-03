package usecase

import (
	"context"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/afif0808/bobobox_test/models"
	"github.com/afif0808/bobobox_test/payloads"
	"github.com/afif0808/bobobox_test/pkg/customerrors"
	"github.com/afif0808/bobobox_test/pkg/snowflake"
	"github.com/afif0808/bobobox_test/pkg/structs"
)

type promoRepository interface {
	InsertPromo(ctx context.Context, p *models.Promo) error
	FetchPromos(ctx context.Context) ([]models.Promo, error)
	UpdatePromo(ctx context.Context, p models.Promo, id int64) error
	DeletePromo(ctx context.Context, id int64) error
	GetPromo(ctx context.Context, id int64) (models.Promo, error)
	GetPromoUsedQuotaByDate(ctx context.Context, promoID int64, date string) (int, error)
}

type PromoUsecase struct {
	repo promoRepository
}

func NewPromoUsecase(repo promoRepository) *PromoUsecase {
	uc := PromoUsecase{
		repo: repo,
	}
	return &uc
}

func convertCreatePromoPayloadToPromo(p payloads.CreatePromoPayload) (models.Promo, error) {
	var pr models.Promo
	err := structs.Merge(&pr, p)
	if err != nil {
		return pr, err
	}
	for _, v := range p.BookingDays {
		pr.BookingDays += strconv.Itoa(int(v)) + ","
	}

	for _, v := range p.CheckInDays {
		pr.CheckInDays += strconv.Itoa(int(v)) + ","
	}
	pr.CheckInDays = strings.TrimSuffix(pr.CheckInDays, ",")
	pr.BookingDays = strings.TrimSuffix(pr.BookingDays, ",")

	return pr, nil
}

func convertUpdatePromoPayloadToPromo(p payloads.UpdatePromoPayload) (models.Promo, error) {
	var pr models.Promo
	err := structs.Merge(&pr, p)
	if err != nil {
		return pr, err
	}
	for _, v := range p.BookingDays {
		pr.BookingDays += strconv.Itoa(int(v)) + ","
	}

	for _, v := range p.CheckInDays {
		pr.CheckInDays += strconv.Itoa(int(v)) + ","
	}
	pr.CheckInDays = strings.TrimSuffix(pr.CheckInDays, ",")
	pr.BookingDays = strings.TrimSuffix(pr.BookingDays, ",")

	return pr, nil
}

func (uc *PromoUsecase) CreatePromo(ctx context.Context, p payloads.CreatePromoPayload) (models.Promo, error) {
	pr, err := convertCreatePromoPayloadToPromo(p)
	if err != nil {
		return pr, err
	}

	pr.ID, err = snowflake.GenerateID()
	if err != nil {
		return pr, err
	}

	err = uc.repo.InsertPromo(ctx, &pr)
	if err != nil {
		return pr, err
	}

	return pr, nil
}

func (uc *PromoUsecase) GetPromoList(ctx context.Context) ([]models.Promo, error) {
	prs, err := uc.repo.FetchPromos(ctx)
	if err != nil {
		return nil, err
	}
	return prs, nil
}

func (uc *PromoUsecase) UpdatePromo(ctx context.Context, p payloads.UpdatePromoPayload, id int64) (models.Promo, error) {
	pr, err := convertUpdatePromoPayloadToPromo(p)
	if err != nil {
		return pr, err
	}
	err = uc.repo.UpdatePromo(ctx, pr, id)
	if err != nil {
		return pr, err
	}
	return pr, nil
}

func (uc *PromoUsecase) DeletePromo(ctx context.Context, id int64) error {
	return uc.repo.DeletePromo(ctx, id)
}

func (uc *PromoUsecase) GetPromo(ctx context.Context, id int64) (models.Promo, error) {
	return uc.repo.GetPromo(ctx, id)
}

func isInHourRange(hourRange string) bool {
	split := strings.Split(hourRange, "-")
	from, until := split[0], split[1]
	timeNow := time.Now().Format("03:04:05")
	if from > timeNow || timeNow > until {
		return false
	}
	return true
}

func calculateTotalPromo(pr models.Promo, roomPrice float64) float64 {
	if pr.Type == "amount" {
		return roomPrice - pr.Value
	}
	// type percentage
	return (pr.Value / 100) * roomPrice
}

func (uc *PromoUsecase) doesPromoRequirementMeet(ctx context.Context, pr models.Promo, p payloads.CheckPromoPayload) (bool, error) {
	todayPromoUsedQuota, err := uc.repo.GetPromoUsedQuotaByDate(ctx, p.PromoID, time.Now().Format("2006-01-02"))
	if err != nil {
		return false, err
	}

	day := time.Now().Weekday()
	isOnBookingDay := strings.Contains(pr.BookingDays, strconv.Itoa(int(day)))
	doesRoomMinimumMeet := len(p.Rooms) >= pr.MinimumNight
	isQuotaAvailable := pr.Quota > 0 && todayPromoUsedQuota < pr.DailyMaxQuota

	if !doesRoomMinimumMeet || !isOnBookingDay || !isInHourRange(pr.BookingHour) || !isQuotaAvailable {
		return false, nil
	}
	return true, nil
}

func processPromo(pr models.Promo, p payloads.CheckPromoPayload) (payloads.CheckPromoResultPayload, error) {
	var result payloads.CheckPromoResultPayload
	for i, r := range p.Rooms {
		result.DefaultTotalPrice += r.Price
		result.FinalPrice += r.Price

		result.Rooms = append(result.Rooms, p.Rooms[i])
		checkInDate, err := time.Parse("2006-01-02", r.CheckInDate)
		if err != nil {
			return result, err
		}
		day := checkInDate.Weekday()
		isOnCheckInDay := strings.Contains(pr.CheckInDays, strconv.Itoa(int(day)))
		if !isOnCheckInDay {
			continue
		}

		// count nights - start
		from := checkInDate
		until, err := time.Parse("2006-01-02", r.CheckOutDate)
		if err != nil {
			return result, err
		}
		nights := int(until.Sub(from).Hours() / 24)

		doesNightsRequirementMeet := nights >= pr.MinimumNight
		if !doesNightsRequirementMeet {
			continue
		}
		result.Rooms[i].Price = calculateTotalPromo(pr, result.Rooms[i].Price)
		promoValue := r.Price - result.Rooms[i].Price
		result.TotalPromo += promoValue
	}
	result.FinalPrice -= result.TotalPromo

	return result, nil

}

func (uc *PromoUsecase) CheckPromo(ctx context.Context, p payloads.CheckPromoPayload) (payloads.CheckPromoResultPayload, error) {
	// Get promo
	// Check if promo requirement meets
	// calculate new price and final price
	var result payloads.CheckPromoResultPayload
	pr, err := uc.repo.GetPromo(ctx, p.PromoID)
	if err != nil {
		return result, nil
	}
	ok, err := uc.doesPromoRequirementMeet(ctx, pr, p)
	if err != nil {
		return result, err
	}
	if !ok {
		err = customerrors.NewCustomError(
			"unprocessable",
			errors.New("payload does not meet promo criteria"),
			customerrors.ErrTypeUnprocessable,
		)
		return result, err
	}
	result, err = processPromo(pr, p)
	return result, err
}
