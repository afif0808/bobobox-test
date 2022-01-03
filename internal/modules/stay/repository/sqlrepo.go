package repository

import (
	"context"
	"log"

	"github.com/afif0808/bobobox_test/models"

	"github.com/jmoiron/sqlx"
)

const (
	tableName = "stays"
)

type StaySQLRepo struct {
	readDB, writeDB *sqlx.DB
}

func NewStaySQLRepo(readDB, writeDB *sqlx.DB) *StaySQLRepo {
	repo := StaySQLRepo{readDB: readDB, writeDB: writeDB}
	return &repo
}

func (repo *StaySQLRepo) fillStayDates(ctx context.Context, ss []models.Stay) error {
	var err error
	for i, v := range ss {
		err = repo.readDB.SelectContext(ctx, &ss[i].Dates, "SELECT * FROM stay_dates WHERE stay_id = ?", v.ID)
		if err != nil {
			break
		}
	}
	return err
}

func (repo *StaySQLRepo) fillReservations(ctx context.Context, ss []models.Stay) error {
	var err error
	for i, v := range ss {
		err = repo.readDB.GetContext(ctx, &ss[i].Reservation, "SELECT * FROM reservations WHERE id = ?", v.ReservationID)
		if err != nil {
			break
		}
	}
	return err
}

func (repo *StaySQLRepo) FetchReservationStays(ctx context.Context, id int64) ([]models.Stay, error) {
	query := "SELECT * FROM " + tableName + " WHERE reservation_id = ? "
	var ss []models.Stay
	err := repo.readDB.SelectContext(ctx, &ss, query, id)
	if err != nil {
		return nil, err
	}
	log.Println(query, id)
	err = repo.fillReservations(ctx, ss)
	if err != nil {
		return nil, err
	}

	err = repo.fillStayDates(ctx, ss)
	if err != nil {
		return nil, err
	}

	return ss, nil
}
