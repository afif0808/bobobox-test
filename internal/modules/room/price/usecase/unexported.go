package usecase

import (
	"context"
	"time"

	"github.com/afif0808/bobobox_test/models"
	"github.com/afif0808/bobobox_test/pkg/snowflake"
)

func parseDates(date, untilDate string) (time.Time, time.Time, error) {
	var parsedDate time.Time
	var parsedUntilDate time.Time
	parsedDate, err := time.Parse("2006-1-2", date)
	if err != nil {
		return parsedDate, parsedUntilDate, err
	}
	parsedUntilDate, err = time.Parse("2006-1-2", untilDate)
	if err != nil {
		return parsedDate, parsedUntilDate, err
	}

	return parsedDate, parsedUntilDate, nil
}

func (uc *RoomPriceUsecase) createOneRoomPrice(ctx context.Context, rp models.RoomPrice) (models.RoomPrice, error) {
	var err error
	rp.ID, err = snowflake.GenerateID()
	if err != nil {
		return rp, err
	}
	err = uc.rpRepo.InsertRoomPrice(ctx, &rp)
	if err != nil {
		return rp, err
	}
	return rp, nil
}

func (uc *RoomPriceUsecase) createManyRoomPrice(ctx context.Context, rp models.RoomPrice, untilDate string) error {
	date, until, err := parseDates(rp.Date, untilDate)
	if err != nil {
		return err
	}
	rp.ID, err = snowflake.GenerateID()
	if err != nil {
		return err
	}
	rps := []models.RoomPrice{rp}
	for date.Before(until) {
		time.Sleep(time.Millisecond) // somehow snowflake needs some time to generate new unique id
		date = date.Add(time.Hour * 24)
		rp.Date = date.Format("2006-01-02")
		if err != nil {
			return err
		}
		rp.ID, err = snowflake.GenerateID()
		rps = append(rps, rp)
	}
	return uc.rpRepo.InsertManyRoomPrice(ctx, rps)

}

func (uc *RoomPriceUsecase) checkRoomType(ctx context.Context, id int64) error {
	_, err := uc.rtRepo.GetRoomType(ctx, id)
	return err
}
