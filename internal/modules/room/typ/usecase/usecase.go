package usecase

import (
	"context"

	"github.com/afif0808/bobobox_test/models"
	"github.com/afif0808/bobobox_test/payloads"
	"github.com/afif0808/bobobox_test/pkg/snowflake"
	"github.com/afif0808/bobobox_test/pkg/structs"
)

type roomTypeRepository interface {
	InsertRoomType(ctx context.Context, rt *models.RoomType) error
	FetchRoomTypes(ctx context.Context) ([]models.RoomType, error)
	UpdateRoomType(ctx context.Context, rt models.RoomType, id int64) error
	DeleteRoomType(ctx context.Context, id int64) error
	GetRoomType(ctx context.Context, id int64) (models.RoomType, error)
}

type RoomTypeUsecase struct {
	repo roomTypeRepository
}

func NewRoomTypeUsecase(repo roomTypeRepository) *RoomTypeUsecase {
	uc := RoomTypeUsecase{
		repo: repo,
	}
	return &uc
}

func (uc *RoomTypeUsecase) CreateRoomType(ctx context.Context, p payloads.CreateRoomTypePayload) (models.RoomType, error) {
	var rt models.RoomType
	err := structs.Merge(&rt, p)
	if err != nil {
		return rt, err
	}

	rt.ID, err = snowflake.GenerateID()
	if err != nil {
		return rt, err
	}

	err = uc.repo.InsertRoomType(ctx, &rt)
	if err != nil {
		return rt, err
	}

	return rt, nil
}

func (uc *RoomTypeUsecase) GetRoomTypeList(ctx context.Context) ([]models.RoomType, error) {
	rts, err := uc.repo.FetchRoomTypes(ctx)
	if err != nil {
		return nil, err
	}
	return rts, nil
}

func (uc *RoomTypeUsecase) UpdateRoomType(ctx context.Context, p payloads.UpdateRoomTypePayload, id int64) (models.RoomType, error) {
	var rt models.RoomType
	err := structs.Merge(&rt, p)
	if err != nil {
		return rt, err
	}
	err = uc.repo.UpdateRoomType(ctx, rt, id)
	if err != nil {
		return rt, err
	}
	return rt, nil
}

func (uc *RoomTypeUsecase) DeleteRoomType(ctx context.Context, id int64) error {
	return uc.repo.DeleteRoomType(ctx, id)
}

func (uc *RoomTypeUsecase) GetRoomType(ctx context.Context, id int64) (models.RoomType, error) {
	return uc.repo.GetRoomType(ctx, id)
}
