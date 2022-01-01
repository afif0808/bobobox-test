package repository

import (
	"context"
	"database/sql"

	"github.com/afif0808/bobobox_test/models"
	"github.com/afif0808/bobobox_test/pkg/customerrors"
	"github.com/afif0808/bobobox_test/pkg/sqls"

	"github.com/jmoiron/sqlx"
)

const (
	tableName = "reservations"
)

type ReservationSQLRepo struct {
	readDB, writeDB *sqlx.DB
}

func NewReservationSQLRepo(readDB, writeDB *sqlx.DB) *ReservationSQLRepo {
	repo := ReservationSQLRepo{readDB: readDB, writeDB: writeDB}
	return &repo
}

func (repo *ReservationSQLRepo) InsertReservation(ctx context.Context, re *models.Reservation) error {
	tx := repo.writeDB.MustBegin()
	defer tx.Commit()

	err := repo.insertReservation(ctx, tx, re)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = repo.insertStays(ctx, tx, re)
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (repo *ReservationSQLRepo) FetchReservations(ctx context.Context) ([]models.Reservation, error) {
	query := "SELECT * FROM " + tableName
	var res []models.Reservation
	err := repo.readDB.SelectContext(ctx, &res, query)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (repo *ReservationSQLRepo) DeleteReservation(ctx context.Context, id int64) error {
	err := sqls.Delete(ctx, repo.writeDB, tableName, id)
	if err == sql.ErrNoRows {
		return customerrors.NewCustomError("reotel is not found", err, customerrors.ErrTypeNotFound)
	}
	return err
}

func (repo *ReservationSQLRepo) GetReservation(ctx context.Context, id int64) (models.Reservation, error) {
	query := "SELECT * FROM " + tableName + " WHERE id = ?"
	var re models.Reservation
	err := repo.readDB.GetContext(ctx, &re, query, id)
	if err == sql.ErrNoRows {
		return re, customerrors.NewCustomError("reotel is not found", err, customerrors.ErrTypeNotFound)
	}
	return re, err
}
