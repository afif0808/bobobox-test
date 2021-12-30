package usecase

import (
	"context"

	"github.com/afif0808/bobobox_test/models"
	"github.com/afif0808/bobobox_test/payloads"
	"github.com/afif0808/bobobox_test/pkg/snowflake"
	"github.com/afif0808/bobobox_test/pkg/structs"
)

type hotelRepository interface {
	InsertHotel(ctx context.Context, h *models.Hotel) error
	FetchHotels(ctx context.Context) ([]models.Hotel, error)
	UpdateHotel(ctx context.Context, h models.Hotel, id int64) error
	DeleteHotel(ctx context.Context, id int64) error
}

type HotelUsecase struct {
	repo hotelRepository
}

func NewHotelUsecase(repo hotelRepository) *HotelUsecase {
	uc := HotelUsecase{
		repo: repo,
	}
	return &uc
}

func (uc *HotelUsecase) CreateHotel(ctx context.Context, p payloads.CreateHotelPayload) (models.Hotel, error) {
	var h models.Hotel
	err := structs.Merge(&h, p)
	if err != nil {
		return h, err
	}

	h.ID, err = snowflake.GenerateID()
	if err != nil {
		return h, err
	}

	err = uc.repo.InsertHotel(ctx, &h)
	if err != nil {
		return h, err
	}

	return h, nil
}

func (uc *HotelUsecase) GetHotelList(ctx context.Context) ([]models.Hotel, error) {
	hs, err := uc.repo.FetchHotels(ctx)
	if err != nil {
		return nil, err
	}
	return hs, nil
}

func (uc *HotelUsecase) UpdateHotel(ctx context.Context, p payloads.UpdateHotelPayload, id int64) (models.Hotel, error) {
	var h models.Hotel
	err := structs.Merge(&h, p)
	if err != nil {
		return h, err
	}
	err = uc.repo.UpdateHotel(ctx, h, id)
	if err != nil {
		return h, err
	}
	return h, nil
}

func (uc *HotelUsecase) DeleteHotel(ctx context.Context, id int64) error {
	return uc.repo.DeleteHotel(ctx, id)
}
