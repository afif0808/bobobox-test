package models

import (
	"github.com/afif0808/bobobox_test/payloads"
	"github.com/afif0808/bobobox_test/pkg/structs"
)

type Room struct {
	ID          int64    `db:"id" action:"create" merge:"id" gorm:"primaryKey;autoIncrement:false"`
	Number      string   `db:"number" action:"create,update" merge:"number"`
	IsInService bool     `db:"is_in_service" action:"create" merge:"is_in_service"`
	RoomTypeID  int64    `db:"room_type_id" action:"create,update" merge:"room_type_id"`
	HotelID     int64    `db:"hotel_id" action:"create" merge:"hotel_id"`
	RoomType    RoomType `gorm:"-"`
	Hotel       Hotel    `gorm:"-"`
}

func (r Room) ToPayload() payloads.RoomPayload {
	var p payloads.RoomPayload
	structs.Merge(&p, r)
	p.RoomType = r.RoomType.ToPayload()
	p.Hotel = r.Hotel.ToPayload()
	return p
}
