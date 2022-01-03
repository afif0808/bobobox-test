package repository

import (
	"context"
	"errors"
	"time"

	"github.com/afif0808/bobobox_test/models"
	"github.com/afif0808/bobobox_test/pkg/customerrors"
	"github.com/afif0808/bobobox_test/pkg/snowflake"
	"github.com/afif0808/bobobox_test/pkg/sqls"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func parseDates(checkInDate, checkOutDate string) (time.Time, time.Time, error) {
	var parsedDate time.Time
	var parsedUntilDate time.Time
	parsedDate, err := time.Parse("2006-1-2", checkInDate)
	if err != nil {
		return parsedDate, parsedUntilDate, err
	}
	parsedUntilDate, err = time.Parse("2006-1-2", checkOutDate)
	if err != nil {
		return parsedDate, parsedUntilDate, err
	}

	return parsedDate, parsedUntilDate, nil
}

func (repo *ReservationSQLRepo) insertReservation(ctx context.Context, tx *sqlx.Tx, re *models.Reservation) error {
	query, args := sqls.GenerateInsertQuery(tableName, re)
	_, err := tx.ExecContext(ctx, query, args...)
	return err
}
func (repo *ReservationSQLRepo) insertStayDates(ctx context.Context, tx *sqlx.Tx, s models.Stay, checkInDate, checkOutDate string) error {
	date, untilDate, err := parseDates(checkInDate, checkOutDate)
	if err != nil {
		return err
	}
	for date.Before(untilDate) || date.Equal(untilDate) {
		time.Sleep(time.Millisecond)
		sd := models.StayDate{
			RoomID: s.RoomID,
			StayID: s.ID,
			Date:   date.Format("2006-01-02"),
		}
		sd.ID, err = snowflake.GenerateID()
		if err != nil {
			break
		}

		query, args := sqls.GenerateInsertQuery("stay_dates", sd)
		_, err = tx.ExecContext(ctx, query, args...)
		if err != nil {
			break
		}

		date = date.Add(time.Hour * 24)
	}
	return err
}

func (repo *ReservationSQLRepo) insertStays(ctx context.Context, tx *sqlx.Tx, re *models.Reservation) error {
	var err error
	for _, s := range re.Stays {
		s.ID, err = snowflake.GenerateID()
		if err != nil {
			break
		}
		s.ReservationID = re.ID
		query, args := sqls.GenerateInsertQuery("stays", s)
		_, err = tx.ExecContext(ctx, query, args...)
		if err != nil {
			break
		}

		err = repo.insertStayDates(ctx, tx, s, re.CheckInDate, re.CheckOutDate)
		if mysqlErr, ok := err.(*mysql.MySQLError); err != nil && ok && mysqlErr.Number == 1062 {
			err = customerrors.NewCustomError(
				"already reserved",
				errors.New("room is already reserved on that date"),
				customerrors.ErrTypeConflict,
			)
			break
		} else if err != nil {
			return err
		}
	}

	return err
}

func (repo *ReservationSQLRepo) usePromo(ctx context.Context, tx *sqlx.Tx, pru *models.PromoUse) error {
	if pru == nil {
		return nil
	}
	query, args := sqls.GenerateInsertQuery("promo_uses", pru)
	_, err := tx.ExecContext(ctx, query, args...)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.ExecContext(ctx, "UPDATE promos SET quota = quota - 1")
	return err
}
