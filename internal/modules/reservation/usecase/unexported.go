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
	"github.com/afif0808/bobobox_test/pkg/structs"
)

func convertCreateStayPayloadToStayModel(ps []payloads.CreateStayPayload) ([]models.Stay, error) {
	var err error
	ss := make([]models.Stay, len(ps))
	for i := range ps {
		err = structs.Merge(&ss[i], ps[i])
		if err != nil {
			return nil, err
		}
	}
	return ss, err
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

func (uc *ReservationUsecase) checkPromo(ctx context.Context, p payloads.CreateReservationPayload) error {
	if p.PromoID == 0 {
		return nil
	}
	pr, err := uc.prRepo.GetPromo(ctx, p.PromoID)
	if err != nil {
		return err
	}
	todayPromoUsedQuota, err := uc.prRepo.GetPromoUsedQuotaByDate(ctx, p.PromoID, time.Now().Format("2006-01-02"))
	if err != nil {
		return err
	}

	day := time.Now().Weekday()
	isOnBookingDay := strings.Contains(pr.BookingDays, strconv.Itoa(int(day)))
	doesRoomMinimumMeet := len(p.Stays) >= pr.MinimumNight
	isQuotaAvailable := pr.Quota > 0 && todayPromoUsedQuota < pr.DailyMaxQuota

	if !doesRoomMinimumMeet || !isOnBookingDay || !isInHourRange(pr.BookingHour) || !isQuotaAvailable {
		err = customerrors.NewCustomError(
			"unprocessable",
			errors.New("payload does not meet promo criteria"),
			customerrors.ErrTypeUnprocessable,
		)
		return err
	}
	return nil
}
