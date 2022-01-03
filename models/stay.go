package models

import (
	"log"

	"github.com/afif0808/bobobox_test/payloads"
	"github.com/afif0808/bobobox_test/pkg/structs"
)

type Stay struct {
	ID            int64       `db:"id" action:"create" merge:"id" gorm:"primaryKey;autoIncrement:false"`
	RoomID        int64       `db:"room_id" action:"create" merge:"room_id"`
	GuestName     string      `db:"guest_name" action:"create" merge:"guest_name"`
	ReservationID int64       `db:"reservation_id" action:"create" merge:"reservation_id"`
	Reservation   Reservation `gorm:"constraint:onDelete:CASCADE"`
	Dates         []StayDate  `gorm:"-"`
}

type StayDate struct {
	ID     int64  `db:"id" action:"create" gorm:"primaryKey;autoIncrement:false"`
	RoomID int64  `db:"room_id" action:"create" gorm:"uniqueIndex:stay_room_date_idx"`
	Date   string `db:"date" action:"create" gorm:"type:date;uniqueIndex:stay_room_date_idx"`
	StayID int64  `db:"stay_id" action:"create"  gorm:"constraint:onDelete:CASCADE"`
	Stay   Stay   `gorm:"constraint:onDelete:CASCADE"`
}

func (s Stay) ToPayload() payloads.StayPayload {
	var p payloads.StayPayload
	structs.Merge(&p, s)

	p.Reservation = s.Reservation.ToPayload()
	p.Dates = make([]payloads.StayDatePayload, len(s.Dates))
	for i := range p.Dates {
		p.Dates[i].Date = s.Dates[i].Date
	}
	log.Println(s.Dates, p.Dates)

	return p
}
