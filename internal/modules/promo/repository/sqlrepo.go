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
	tableName = "promos"
)

type PromoSQLRepo struct {
	readDB, writeDB *sqlx.DB
}

func NewPromoSQLRepo(readDB, writeDB *sqlx.DB) *PromoSQLRepo {
	repo := PromoSQLRepo{readDB: readDB, writeDB: writeDB}
	return &repo
}

func (repo *PromoSQLRepo) InsertPromo(ctx context.Context, pr *models.Promo) error {
	err := sqls.Insert(ctx, repo.writeDB, tableName, pr)
	if err != nil {
		return err
	}
	return nil
}
func (repo *PromoSQLRepo) UpdatePromo(ctx context.Context, pr models.Promo, id int64) error {
	err := sqls.Update(ctx, repo.writeDB, tableName, pr, id)
	if err == sql.ErrNoRows {
		return customerrors.NewCustomError("room type is not found", err, customerrors.ErrTypeNotFound)
	}
	return err

}

func (repo *PromoSQLRepo) FetchPromos(ctx context.Context) ([]models.Promo, error) {
	query := "SELECT * FROM " + tableName
	var prs []models.Promo
	err := repo.readDB.SelectContext(ctx, &prs, query)
	if err != nil {
		return nil, err
	}
	return prs, nil
}

func (repo *PromoSQLRepo) DeletePromo(ctx context.Context, id int64) error {
	err := sqls.Delete(ctx, repo.writeDB, tableName, id)
	if err == sql.ErrNoRows {
		return customerrors.NewCustomError("room type is not found", err, customerrors.ErrTypeNotFound)
	}
	return err
}

func (repo *PromoSQLRepo) GetPromo(ctx context.Context, id int64) (models.Promo, error) {
	query := "SELECT * FROM " + tableName + " WHERE id = ?"
	var pr models.Promo
	err := repo.readDB.GetContext(ctx, &pr, query, id)
	if err == sql.ErrNoRows {
		return pr, customerrors.NewCustomError("room type is not found", err, customerrors.ErrTypeNotFound)
	}
	return pr, err
}

func (repo *PromoSQLRepo) GetPromoUsedQuotaByDate(ctx context.Context, promoID int64, date string) (int, error) {
	query := "SELECT count(*) FROM  promo_uses  WHERE promo_id = ? AND booking_date = ?"
	var count int
	err := repo.readDB.GetContext(ctx, &count, query, promoID, date)
	return count, err
}
