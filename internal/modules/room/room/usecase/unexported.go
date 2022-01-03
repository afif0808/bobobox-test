package usecase

import (
	"context"
	"time"

	"github.com/afif0808/bobobox_test/models"
)

func (uc *RoomUsecase) getRoomTypeAndHotel(ctx context.Context, r *models.Room) error {
	var err error
	r.RoomType, err = uc.rtRepo.GetRoomType(ctx, r.RoomTypeID)
	if err != nil {
		return err
	}
	r.Hotel, err = uc.hRepo.GetHotel(ctx, r.HotelID)
	return err
}

func (uc *RoomUsecase) checkRoomType(ctx context.Context, id int64) error {
	_, err := uc.rtRepo.GetRoomType(ctx, id)
	return err
}

// fillMissingPrice fill undefined room price at given date  if any
func (uc *RoomUsecase) fillMissingPrice(ctx context.Context, rps []models.RoomPrice, defaultPrice float64, checkInDate, checkOutDate string) ([]models.RoomPrice, error) {
	from, err := time.Parse("2006-1-2", checkInDate)
	if err != nil {
		return rps, err
	}
	until, err := time.Parse("2006-1-2", checkOutDate)
	if err != nil {
		return rps, err
	}
	dayCount := until.Sub(from).Hours()/24 + 1

	if len(rps) == int(dayCount) {
		return rps, nil
	}

	result := make([]models.RoomPrice, int(dayCount))

	for i, j := 0, 0; from.Before(until) || from.Equal(until); j++ {
		if len(rps) != 0 && i < len(rps) && from.Format("2006-01-02") == rps[i].Date {
			result[j] = rps[i]
			i++
		} else {
			rp := models.RoomPrice{
				Date:  from.Format("2006-01-02"),
				Price: defaultPrice,
			}
			result[j] = rp
		}
		from = from.Add(24 * time.Hour)
	}

	return result, nil
}

func (uc *RoomUsecase) getAvailableRoomPrices(ctx context.Context, checkInDate, checkOutDate string, roomTypeID int64) (rps []models.RoomPrice, totalPrice float64, err error) {
	rt, err := uc.rtRepo.GetRoomType(ctx, roomTypeID)
	if err != nil {
		return
	}
	rps, err = uc.rpRepo.FetchRoomTypePricesByDateRange(ctx, roomTypeID, checkInDate, checkOutDate)
	if err != nil {
		return
	}
	rps, err = uc.fillMissingPrice(ctx, rps, rt.DefaultPrice, checkInDate, checkOutDate)
	if err != nil {
		return
	}
	totalPrice = uc.calculateTotalPrice(rps)

	return
}

func (uc *RoomUsecase) calculateTotalPrice(rps []models.RoomPrice) float64 {
	var total float64
	for _, v := range rps {
		total += v.Price
	}
	return total
}
