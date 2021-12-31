package sqls

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

func Update(ctx context.Context, db *sqlx.DB, table string, data, id interface{}) error {
	err := Check(ctx, db, table, id)
	if err != nil {
		return err
	}
	query, args := GenerateUpdateByIDQuery(table, data, id)

	_, err = db.ExecContext(ctx, query, args...)
	return err
}

func Delete(ctx context.Context, db *sqlx.DB, table string, id interface{}) error {
	res, err := db.ExecContext(ctx, "DELETE FROM "+table+" WHERE id = ?", id)
	if err != nil {
		return err
	}
	affected, err := res.RowsAffected()
	if affected < 1 {
		return sql.ErrNoRows
	}
	return err
}

func Insert(ctx context.Context, db *sqlx.DB, table string, data interface{}) error {
	query, args := GenerateInsertQuery(table, data)
	_, err := db.ExecContext(ctx, query, args...)
	return err
}

func Check(ctx context.Context, db *sqlx.DB, table string, id interface{}) error {
	row := db.QueryRowContext(ctx, "SELECT NULL FROM "+table+" WHERE id = ?", id)
	var any interface{}
	return row.Scan(&any)
}
