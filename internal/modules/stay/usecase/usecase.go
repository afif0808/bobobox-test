package usecase

import (
	"context"

	"github.com/afif0808/bobobox_test/models"
)

type stayRepository interface {
	FetchReservationStays(ctx context.Context, id int64) ([]models.Stay, error)
}

type StayUsecase struct {
	repo stayRepository
}

func NewStayUsecase(repo stayRepository) *StayUsecase {
	uc := StayUsecase{
		repo: repo,
	}
	return &uc
}

func (uc *StayUsecase) GetReservationStayList(ctx context.Context, id int64) ([]models.Stay, error) {
	ss, err := uc.repo.FetchReservationStays(ctx, id)
	if err != nil {
		return nil, err
	}
	return ss, nil
}
