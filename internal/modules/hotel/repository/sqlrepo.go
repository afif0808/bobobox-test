package repository

import (
	"context"

	"github.com/afif0808/bobobox_test/models"
	"github.com/afif0808/bobobox_test/pkg/sqls"

	"github.com/jmoiron/sqlx"
)

const (
	tableName = "hotels"
)

type SQLRepo struct {
	readDB, writeDB *sqlx.DB
}

func NewHotelSQLRepo(readDB, writeDB *sqlx.DB) *SQLRepo {
	repo := SQLRepo{readDB: readDB, writeDB: writeDB}
	return &repo
}

func (repo *SQLRepo) InsertHotel(ctx context.Context, h *models.Hotel) error {
	query, args := sqls.GenerateInsertQuery(tableName, *h)
	_, err := repo.writeDB.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}
	return nil
}
func (repo *SQLRepo) UpdateHotel(ctx context.Context, h models.Hotel, id int64) error {
	query, args := sqls.GenerateUpdateByIDQuery(tableName, h, id)
	_, err := repo.writeDB.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}
	return nil
}

func (repo *SQLRepo) FetchHotels(ctx context.Context) ([]models.Hotel, error) {
	query := "SELECT * FROM " + tableName
	var hs []models.Hotel
	err := repo.readDB.SelectContext(ctx, &hs, query)
	if err != nil {
		return nil, err
	}
	return hs, nil

}
