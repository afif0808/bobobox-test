package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/afif0808/bobobox_test/models"
	"github.com/afif0808/bobobox_test/pkg/customerrors"
	"github.com/afif0808/bobobox_test/pkg/sqls"

	"github.com/jmoiron/sqlx"
)

const (
	tableName = "hotels"
)

type HotelSQLRepo struct {
	readDB, writeDB *sqlx.DB
}

func NewHotelSQLRepo(readDB, writeDB *sqlx.DB) *HotelSQLRepo {
	repo := HotelSQLRepo{readDB: readDB, writeDB: writeDB}
	return &repo
}

func (repo *HotelSQLRepo) InsertHotel(ctx context.Context, h *models.Hotel) error {
	query, args := sqls.GenerateInsertQuery(tableName, *h)
	_, err := repo.writeDB.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}
	return nil
}
func (repo *HotelSQLRepo) UpdateHotel(ctx context.Context, h models.Hotel, id int64) error {
	query, args := sqls.GenerateUpdateByIDQuery(tableName, h, id)
	res, err := repo.writeDB.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affected < 1 {
		return customerrors.NewCustomError("hotel is not found", errors.New("no rows is affected"), customerrors.ErrTypeNotFound)
	}

	return nil
}

func (repo *HotelSQLRepo) FetchHotels(ctx context.Context) ([]models.Hotel, error) {
	query := "SELECT * FROM " + tableName
	var hs []models.Hotel
	err := repo.readDB.SelectContext(ctx, &hs, query)
	if err != nil {
		return nil, err
	}
	return hs, nil
}

func (repo *HotelSQLRepo) DeleteHotel(ctx context.Context, id int64) error {
	query := "DELETE FROM " + tableName + " WHERE id = ?"
	res, err := repo.writeDB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affected < 1 {
		return customerrors.NewCustomError("hotel is not found", errors.New("no rows is affected"), customerrors.ErrTypeNotFound)
	}

	return nil
}

func (repo *HotelSQLRepo) GetHotel(ctx context.Context, id int64) (models.Hotel, error) {
	query := "SELECT * FROM " + tableName + " WHERE id = ?"
	var h models.Hotel
	err := repo.readDB.GetContext(ctx, &h, query, id)
	if err != nil && err == sql.ErrNoRows {
		return h, customerrors.NewCustomError("hotel is not found", err, customerrors.ErrTypeNotFound)
	} else if err != nil {
		return h, err
	}
	return h, nil
}
