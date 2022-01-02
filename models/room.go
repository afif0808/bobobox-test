package models

import (
	"github.com/afif0808/bobobox_test/payloads"
	"github.com/afif0808/bobobox_test/pkg/structs"
)

type Room struct {
	ID          int64    `db:"id" action:"create" merge:"id" gorm:"primaryKey;autoIncrement:false"`
	Number      string   `db:"number" action:"create,update" merge:"number"`
	IsInService bool     `db:"is_in_service" action:"create,update" merge:"is_in_service"`
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

type AvailableRoom struct {
	ID        int64  `db:"id" merge:"id"`
	Number    string `db:"number" merge:"number"`
	Prices    []RoomPrice
	HotelID   int64  `db:"hotel_id" merge:"hotel_id"`
	HotelName string `db:"hotel_name" merge:"hotel_name"`
}

func (ar AvailableRoom) ToPayload() payloads.AvailableRoomPayload {
	var p payloads.AvailableRoomPayload
	structs.Merge(&p, ar)
	p.Prices = make([]payloads.AvailableRoomPricePayload, len(ar.Prices))
	for i := range ar.Prices {
		p.Prices[i] = ar.Prices[i].ToAvailableRoomPricePayload()
	}

	return p
}
