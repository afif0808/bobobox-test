package usecase

import (
	"context"

	"github.com/afif0808/bobobox_test/models"
	"github.com/afif0808/bobobox_test/payloads"
	"github.com/afif0808/bobobox_test/pkg/snowflake"
	"github.com/afif0808/bobobox_test/pkg/structs"
)

type roomTypeRepository interface {
	InsertRoomType(ctx context.Context, h *models.RoomType) error
	FetchRoomTypes(ctx context.Context) ([]models.RoomType, error)
	UpdateRoomType(ctx context.Context, h models.RoomType, id int64) error
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
	var h models.RoomType
	err := structs.Merge(&h, p)
	if err != nil {
		return h, err
	}

	h.ID, err = snowflake.GenerateID()
	if err != nil {
		return h, err
	}

	err = uc.repo.InsertRoomType(ctx, &h)
	if err != nil {
		return h, err
	}

	return h, nil
}

func (uc *RoomTypeUsecase) GetRoomTypeList(ctx context.Context) ([]models.RoomType, error) {
	hs, err := uc.repo.FetchRoomTypes(ctx)
	if err != nil {
		return nil, err
	}
	return hs, nil
}

func (uc *RoomTypeUsecase) UpdateRoomType(ctx context.Context, p payloads.UpdateRoomTypePayload, id int64) (models.RoomType, error) {
	var h models.RoomType
	err := structs.Merge(&h, p)
	if err != nil {
		return h, err
	}
	err = uc.repo.UpdateRoomType(ctx, h, id)
	if err != nil {
		return h, err
	}
	return h, nil
}

func (uc *RoomTypeUsecase) DeleteRoomType(ctx context.Context, id int64) error {
	return uc.repo.DeleteRoomType(ctx, id)
}

func (uc *RoomTypeUsecase) GetRoomType(ctx context.Context, id int64) (models.RoomType, error) {
	return uc.repo.GetRoomType(ctx, id)
}
