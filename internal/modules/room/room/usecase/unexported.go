package usecase

import (
	"context"

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
