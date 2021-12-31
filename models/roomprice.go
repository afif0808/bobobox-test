package models

import (
	"github.com/afif0808/bobobox_test/payloads"
	"github.com/afif0808/bobobox_test/pkg/structs"
)

type RoomPrice struct {
	ID         int64   `db:"id" action:"create" gorm:"primaryKey;autoIncrement:false" merge:"id"`
	Date       string  `db:"date" gorm:"uniqueIndex:date_type_idx,length:10" action:"create" merge:"date"`
	Price      float64 `db:"price" action:"create,update" merge:"price"`
	RoomTypeID int64   `db:"room_type_id" gorm:"uniqueIndex:date_type_idx" action:"create" merge:"room_type_id"`
}

func (rp RoomPrice) ToPayload() payloads.RoomPricePayload {
	var p payloads.RoomPricePayload
	structs.Merge(&p, rp)
	return p

}
