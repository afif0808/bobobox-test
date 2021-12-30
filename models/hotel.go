package models

import (
	"github.com/afif0808/bobobox_test/payloads"
	"github.com/afif0808/bobobox_test/pkg/structs"
)

type Hotel struct {
	ID      int64  `db:"id" action:"create" merge:"id" gorm:"primaryKey;autoIncrement:false"`
	Name    string `db:"name" action:"create,update" merge:"name"`
	Address string `db:"address" action:"create,update" merge:"address"`
}

func (h Hotel) ToPayload() payloads.HotelPayload {
	var p payloads.HotelPayload
	structs.Merge(&p, h)
	return p
}
