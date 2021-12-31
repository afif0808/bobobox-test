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
	tableName = "room_types"
)

type RoomTypeSQLRepo struct {
	readDB, writeDB *sqlx.DB
}

func NewRoomTypeSQLRepo(readDB, writeDB *sqlx.DB) *RoomTypeSQLRepo {
	repo := RoomTypeSQLRepo{readDB: readDB, writeDB: writeDB}
	return &repo
}

func (repo *RoomTypeSQLRepo) InsertRoomType(ctx context.Context, rt *models.RoomType) error {
	err := sqls.Insert(ctx, repo.writeDB, tableName, rt)
	if err != nil {
		return err
	}
	return nil
}
func (repo *RoomTypeSQLRepo) UpdateRoomType(ctx context.Context, rt models.RoomType, id int64) error {
	err := sqls.Update(ctx, repo.writeDB, tableName, rt, id)
	if err == sql.ErrNoRows {
		return customerrors.NewCustomError("room type is not found", err, customerrors.ErrTypeNotFound)
	}
	return err

}

func (repo *RoomTypeSQLRepo) FetchRoomTypes(ctx context.Context) ([]models.RoomType, error) {
	query := "SELECT * FROM " + tableName
	var hs []models.RoomType
	err := repo.readDB.SelectContext(ctx, &hs, query)
	if err != nil {
		return nil, err
	}
	return hs, nil
}

func (repo *RoomTypeSQLRepo) DeleteRoomType(ctx context.Context, id int64) error {
	err := sqls.Delete(ctx, repo.writeDB, tableName, id)
	if err == sql.ErrNoRows {
		return customerrors.NewCustomError("room type is not found", err, customerrors.ErrTypeNotFound)
	}
	return err
}

func (repo *RoomTypeSQLRepo) GetRoomType(ctx context.Context, id int64) (models.RoomType, error) {
	query := "SELECT * FROM " + tableName + " WHERE id = ?"
	var h models.RoomType
	err := repo.readDB.GetContext(ctx, &h, query, id)
	if err == sql.ErrNoRows {
		return h, customerrors.NewCustomError("room type is not found", err, customerrors.ErrTypeNotFound)
	}
	return h, err
}
