package usecase

import (
	"context"

	"github.com/afif0808/bobobox_test/models"
	"github.com/afif0808/bobobox_test/payloads"
	"github.com/afif0808/bobobox_test/pkg/snowflake"
	"github.com/afif0808/bobobox_test/pkg/structs"
)

type roomRepository interface {
	InsertRoom(ctx context.Context, r *models.Room) error
	FetchRooms(ctx context.Context) ([]models.Room, error)
	UpdateRoom(ctx context.Context, r models.Room, id int64) error
	DeleteRoom(ctx context.Context, id int64) error
	GetRoom(ctx context.Context, id int64) (models.Room, error)
}

type roomTypeRepository interface {
	GetRoomType(ctx context.Context, id int64) (models.RoomType, error)
}

type hotelRepository interface {
	GetHotel(ctx context.Context, id int64) (models.Hotel, error)
}

type RoomUsecase struct {
	rRepo  roomRepository
	rtRepo roomTypeRepository
	hRepo  hotelRepository
}

func NewRoomUsecase(rRepo roomRepository, rtRepo roomTypeRepository, hRepo hotelRepository) *RoomUsecase {
	uc := RoomUsecase{
		rRepo:  rRepo,
		rtRepo: rtRepo,
		hRepo:  hRepo,
	}

	return &uc
}

func (uc *RoomUsecase) CreateRoom(ctx context.Context, p payloads.CreateRoomPayload) (models.Room, error) {
	var r models.Room
	err := structs.Merge(&r, p)
	if err != nil {
		return r, err
	}

	r.ID, err = snowflake.GenerateID()
	if err != nil {
		return r, err
	}

	err = uc.getRoomTypeAndHotel(ctx, &r)
	if err != nil {
		return r, err
	}

	r.IsInService = true
	err = uc.rRepo.InsertRoom(ctx, &r)
	if err != nil {
		return r, err
	}

	return r, nil
}

func (uc *RoomUsecase) GetRoomList(ctx context.Context) ([]models.Room, error) {
	rs, err := uc.rRepo.FetchRooms(ctx)
	if err != nil {
		return nil, err
	}
	return rs, nil
}

func (uc *RoomUsecase) UpdateRoom(ctx context.Context, p payloads.UpdateRoomPayload, id int64) error {
	var r models.Room
	err := structs.Merge(&r, p)
	if err != nil {
		return err
	}

	err = uc.checkRoomType(ctx, r.RoomTypeID)
	if err != nil {
		return err
	}

	err = uc.rRepo.UpdateRoom(ctx, r, id)
	if err != nil {
		return err
	}

	return nil
}

func (uc *RoomUsecase) DeleteRoom(ctx context.Context, id int64) error {
	return uc.rRepo.DeleteRoom(ctx, id)
}

func (uc *RoomUsecase) GetRoom(ctx context.Context, id int64) (models.Room, error) {
	return uc.rRepo.GetRoom(ctx, id)
}
