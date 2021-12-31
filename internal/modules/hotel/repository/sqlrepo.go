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
	err := sqls.Insert(ctx, repo.writeDB, tableName, h)
	if err != nil {
		return err
	}
	return nil
}
func (repo *HotelSQLRepo) UpdateHotel(ctx context.Context, h models.Hotel, id int64) error {
	err := sqls.Update(ctx, repo.writeDB, tableName, h, id)
	if err == sql.ErrNoRows {
		return customerrors.NewCustomError("hotel is not found", err, customerrors.ErrTypeNotFound)
	}
	return err
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
	err := sqls.Delete(ctx, repo.writeDB, tableName, id)
	if err == sql.ErrNoRows {
		return customerrors.NewCustomError("hotel is not found", err, customerrors.ErrTypeNotFound)
	}
	return err
}

func (repo *HotelSQLRepo) GetHotel(ctx context.Context, id int64) (models.Hotel, error) {
	query := "SELECT * FROM " + tableName + " WHERE id = ?"
	var h models.Hotel
	err := repo.readDB.GetContext(ctx, &h, query, id)
	if err == sql.ErrNoRows {
		return h, customerrors.NewCustomError("hotel is not found", err, customerrors.ErrTypeNotFound)
	}
	return h, err
}
