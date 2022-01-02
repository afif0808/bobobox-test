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
	tableName = "rooms"
)

type RoomSQLRepo struct {
	readDB, writeDB *sqlx.DB
}

func NewRoomSQLRepo(readDB, writeDB *sqlx.DB) *RoomSQLRepo {
	repo := RoomSQLRepo{readDB: readDB, writeDB: writeDB}
	return &repo
}

func (repo *RoomSQLRepo) InsertRoom(ctx context.Context, r *models.Room) error {
	err := sqls.Insert(ctx, repo.writeDB, tableName, r)
	if err != nil {
		return err
	}
	return nil
}
func (repo *RoomSQLRepo) UpdateRoom(ctx context.Context, r models.Room, id int64) error {
	err := sqls.Update(ctx, repo.writeDB, tableName, r, id)
	if err == sql.ErrNoRows {
		return customerrors.NewCustomError("room type is not found", err, customerrors.ErrTypeNotFound)
	}
	return err

}

func (repo *RoomSQLRepo) fillManyHotelAndType(ctx context.Context, rs []models.Room) error {
	var err error
	for i, r := range rs {
		err = repo.readDB.GetContext(ctx, &rs[i].Hotel, "SELECT * FROM hotels WHERE id = ? ", r.HotelID)
		if err != nil {
			break
		}
		err = repo.readDB.GetContext(ctx, &rs[i].RoomType, "SELECT * FROM room_types WHERE id = ? ", r.RoomTypeID)
		if err != nil {
			break
		}
	}
	if err == sql.ErrNoRows {
		return customerrors.NewCustomError("room's type or hotel is not found", err, customerrors.ErrTypeNotFound)
	}

	return err
}

func (repo *RoomSQLRepo) fillOneHotelAndType(ctx context.Context, r *models.Room) error {
	err := repo.readDB.GetContext(ctx, &r.Hotel, "SELECT * FROM hotels WHERE id = ? ", r.HotelID)
	if err != nil {
		return err
	}
	err = repo.readDB.GetContext(ctx, &r.RoomType, "SELECT * FROM room_types WHERE id = ? ", r.RoomTypeID)

	return err
}

func (repo *RoomSQLRepo) FetchRooms(ctx context.Context) ([]models.Room, error) {
	query := "SELECT * FROM " + tableName
	var rs []models.Room
	err := repo.readDB.SelectContext(ctx, &rs, query)
	if err != nil {
		return nil, err
	}

	err = repo.fillManyHotelAndType(ctx, rs)

	return rs, err
}

func (repo *RoomSQLRepo) FetchHotelRooms(ctx context.Context, hotelID int64) ([]models.Room, error) {
	query := "SELECT * FROM " + tableName + " WHERE hotel_id = ?"
	var rs []models.Room
	err := repo.readDB.SelectContext(ctx, &rs, query, hotelID)
	if err != nil {
		return nil, err
	}

	err = repo.fillManyHotelAndType(ctx, rs)

	return rs, err
}

func (repo *RoomSQLRepo) DeleteRoom(ctx context.Context, id int64) error {
	err := sqls.Delete(ctx, repo.writeDB, tableName, id)
	if err == sql.ErrNoRows {
		return customerrors.NewCustomError("room type is not found", err, customerrors.ErrTypeNotFound)
	}
	return err
}

func (repo *RoomSQLRepo) GetRoom(ctx context.Context, id int64) (models.Room, error) {
	query := "SELECT * FROM " + tableName + " WHERE id = ?"
	var r models.Room
	err := repo.readDB.GetContext(ctx, &r, query, id)
	if err == sql.ErrNoRows {
		return r, customerrors.NewCustomError("room type is not found", err, customerrors.ErrTypeNotFound)
	}

	err = repo.fillOneHotelAndType(ctx, &r)
	if err == sql.ErrNoRows {
		return r, customerrors.NewCustomError("room's type or hotel is not found", err, customerrors.ErrTypeNotFound)
	}

	return r, err
}

func (repo *RoomSQLRepo) GetAvailableRooms(ctx context.Context, checkInDate, checkOutDate string, roomTypeID int64) ([]models.AvailableRoom, error) {
	query := "SELECT rooms.id, rooms.number , rooms.hotel_id , hotels.name as hotel_name FROM " + tableName + " INNER JOIN hotels ON hotels.id = rooms.hotel_id WHERE room_type_id = ? AND is_in_service = true AND  !EXISTS(SELECT NULL FROM stay_dates WHERE room_id = rooms.id AND stay_dates.date >= ? AND stay_dates.date <= ?)"
	var ars []models.AvailableRoom
	err := repo.readDB.SelectContext(ctx, &ars, query, roomTypeID, checkInDate, checkOutDate)
	if err != nil {
		return nil, err
	}
	return ars, nil
}
