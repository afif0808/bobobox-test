package usecase

import (
	"context"

	"github.com/afif0808/bobobox_test/models"
	"github.com/afif0808/bobobox_test/payloads"
	"github.com/afif0808/bobobox_test/pkg/snowflake"
	"github.com/afif0808/bobobox_test/pkg/structs"
)

type reservationRepository interface {
	InsertReservation(ctx context.Context, re *models.Reservation) error
	FetchReservations(ctx context.Context) ([]models.Reservation, error)
	DeleteReservation(ctx context.Context, id int64) error
	GetReservation(ctx context.Context, id int64) (models.Reservation, error)
}

type ReservationUsecase struct {
	reRepo reservationRepository
}

func NewReservationUsecase(reRepo reservationRepository) *ReservationUsecase {
	uc := ReservationUsecase{
		reRepo: reRepo,
	}
	return &uc
}

func (uc *ReservationUsecase) CreateReservation(ctx context.Context, p payloads.CreateReservationPayload) (models.Reservation, error) {
	var re models.Reservation
	err := structs.Merge(&re, p)
	if err != nil {
		return re, err
	}

	re.ID, err = snowflake.GenerateID()
	if err != nil {
		return re, err
	}
	re.BookedRoomCount = len(p.Stays)

	re.Stays, err = convertCreateStayPayloadToStayModel(p.Stays)
	if err != nil {
		return re, err
	}

	err = uc.reRepo.InsertReservation(ctx, &re)
	if err != nil {
		return re, err
	}

	return re, nil
}

func (uc *ReservationUsecase) GetReservationList(ctx context.Context) ([]models.Reservation, error) {
	res, err := uc.reRepo.FetchReservations(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (uc *ReservationUsecase) DeleteReservation(ctx context.Context, id int64) error {
	return uc.reRepo.DeleteReservation(ctx, id)
}

func (uc *ReservationUsecase) GetReservation(ctx context.Context, id int64) (models.Reservation, error) {
	return uc.reRepo.GetReservation(ctx, id)
}
