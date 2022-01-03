package usecase

import (
	"context"
	"time"

	"github.com/afif0808/bobobox_test/models"
	"github.com/afif0808/bobobox_test/payloads"
	"github.com/afif0808/bobobox_test/pkg/snowflake"
	"github.com/afif0808/bobobox_test/pkg/structs"
)

type reservationRepository interface {
	InsertReservation(ctx context.Context, re *models.Reservation, pru *models.PromoUse) error
	FetchReservations(ctx context.Context) ([]models.Reservation, error)
	DeleteReservation(ctx context.Context, id int64) error
	GetReservation(ctx context.Context, id int64) (models.Reservation, error)
}

type promoRepository interface {
	GetPromo(ctx context.Context, id int64) (models.Promo, error)
	GetPromoUsedQuotaByDate(ctx context.Context, promoID int64, date string) (int, error)
}

type ReservationUsecase struct {
	reRepo reservationRepository
	prRepo promoRepository
}

func NewReservationUsecase(reRepo reservationRepository, prRepo promoRepository) *ReservationUsecase {
	uc := ReservationUsecase{
		reRepo: reRepo,
		prRepo: prRepo,
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

	err = uc.checkPromo(ctx, p)
	if err != nil {
		return re, err
	}
	var pru *models.PromoUse
	// use promo if given
	if p.PromoID != 0 {
		pru = &models.PromoUse{
			ReservationID: re.ID,
			PromoID:       p.PromoID,
			BookingDate:   time.Now().Format("2006-01-02"),
		}
	}
	err = uc.reRepo.InsertReservation(ctx, &re, pru)
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
