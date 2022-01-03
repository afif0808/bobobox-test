package models

import (
	"github.com/afif0808/bobobox_test/payloads"
	"github.com/afif0808/bobobox_test/pkg/structs"
)

type RoomPrice struct {
	ID         int64    `db:"id" action:"create" gorm:"primaryKey;autoIncrement:false" merge:"id"`
	Date       string   `db:"date" gorm:"type:date;uniqueIndex:date_type_idx" action:"create" merge:"date"`
	Price      float64  `db:"price" action:"create,update" merge:"price"`
	RoomTypeID int64    `db:"room_type_id" gorm:"uniqueIndex:date_type_idx" action:"create" merge:"room_type_id"`
	RoomType   RoomType `db:"constraint:onDelete:CASCADE"`
}

func (rp RoomPrice) ToPayload() payloads.RoomPricePayload {
	var p payloads.RoomPricePayload
	structs.Merge(&p, rp)
	return p
}

func (rp RoomPrice) ToAvailableRoomPricePayload() payloads.AvailableRoomPricePayload {
	var p payloads.AvailableRoomPricePayload
	structs.Merge(&p, rp)
	return p
}
