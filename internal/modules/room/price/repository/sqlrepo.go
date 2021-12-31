package repository

import (
	"context"
	"database/sql"

	"github.com/afif0808/bobobox_test/models"
	"github.com/afif0808/bobobox_test/pkg/customerrors"
	"github.com/afif0808/bobobox_test/pkg/sqls"
	"github.com/go-sql-driver/mysql"

	"github.com/jmoiron/sqlx"
)

const (
	tableName = "room_prices"
)

type RoomPriceSQLRepo struct {
	readDB, writeDB *sqlx.DB
}

func NewRoomPriceSQLRepo(readDB, writeDB *sqlx.DB) *RoomPriceSQLRepo {
	repo := RoomPriceSQLRepo{readDB: readDB, writeDB: writeDB}
	return &repo
}

func (repo *RoomPriceSQLRepo) InsertRoomPrice(ctx context.Context, rp *models.RoomPrice) error {
	err := sqls.Insert(ctx, repo.writeDB, tableName, rp)
	if err != nil {
		return err
	}
	return nil
}
func (repo *RoomPriceSQLRepo) UpdateRoomPrice(ctx context.Context, rp models.RoomPrice, id int64) error {
	err := sqls.Update(ctx, repo.writeDB, tableName, rp, id)
	if err == sql.ErrNoRows {
		return customerrors.NewCustomError("room type is not found", err, customerrors.ErrTypeNotFound)
	}
	return err

}

func (repo *RoomPriceSQLRepo) FetchRoomTypePrices(ctx context.Context, id int64) ([]models.RoomPrice, error) {
	query := "SELECT * FROM " + tableName + " WHERE room_type_id = ?"
	var rps []models.RoomPrice
	err := repo.readDB.SelectContext(ctx, &rps, query, id)
	if err != nil {
		return nil, err
	}
	return rps, nil
}

func (repo *RoomPriceSQLRepo) DeleteRoomPrice(ctx context.Context, id int64) error {
	err := sqls.Delete(ctx, repo.writeDB, tableName, id)
	if err == sql.ErrNoRows {
		return customerrors.NewCustomError("room type is not found", err, customerrors.ErrTypeNotFound)
	}
	return err
}

func (repo *RoomPriceSQLRepo) GetRoomPrice(ctx context.Context, id int64) (models.RoomPrice, error) {
	query := "SELECT * FROM " + tableName + " WHERE id = ?"
	var rp models.RoomPrice
	err := repo.readDB.GetContext(ctx, &rp, query, id)
	if err == sql.ErrNoRows {
		return rp, customerrors.NewCustomError("room type is not found", err, customerrors.ErrTypeNotFound)
	}
	return rp, err
}

func (repo *RoomPriceSQLRepo) InsertManyRoomPrice(ctx context.Context, rps []models.RoomPrice) error {
	tx := repo.writeDB.MustBegin().Tx
	defer tx.Commit()
	query, arg := sqls.GenerateInsertQuery(tableName, rps[0])
	args := [][]interface{}{arg}
	for i := 1; i < len(rps); i++ {
		args = append(args, sqls.GenerateArgs(rps[i]))
	}
	var err error
	for i := range rps {
		_, err = tx.ExecContext(ctx, query, args[i]...)
		if err != nil {
			tx.Rollback()
			break
		}

	}

	if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1062 {
		err = customerrors.NewCustomError("duplicate room price:", err, customerrors.ErrTypeConflict)
	}

	return err
}
