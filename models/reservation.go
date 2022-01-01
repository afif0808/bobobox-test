package models

import (
	"github.com/afif0808/bobobox_test/payloads"
	"github.com/afif0808/bobobox_test/pkg/structs"
)

type Reservation struct {
	ID              int64  `db:"id" action:"create" merge:"id" gorm:"primaryKey;autoIncrement:false"`
	CustomerName    string `db:"customer_name" action:"create" merge:"customer_name"`
	BookedRoomCount int    `db:"booked_room_count" action:"create" merge:"booked_room_count"`
	CheckInDate     string `db:"check_in_date" action:"create" merge:"check_in_date" gorm:"type:date"`
	CheckOutDate    string `db:"check_out_date" action:"create" merge:"check_out_date" gorm:"type:date"`
	HotelID         int64  `db:"hotel_id" action:"create" merge:"hotel_id"`
	Stays           []Stay `gorm:"-"`
}

func (re Reservation) ToPayload() payloads.ReservationPayload {
	var p payloads.ReservationPayload
	structs.Merge(&p, re)
	return p
}
