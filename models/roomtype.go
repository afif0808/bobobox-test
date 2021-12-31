package models

import (
	"github.com/afif0808/bobobox_test/payloads"
	"github.com/afif0808/bobobox_test/pkg/structs"
)

type RoomType struct {
	ID   int64  `db:"id" action:"create" gorm:"primaryKey,autoIncrement:false" merge:"id"`
	Name string `db:"name" action:"create,update" merge:"name"`
}

func (rt RoomType) ToPayload() payloads.RoomTypePayload {
	var p payloads.RoomTypePayload
	structs.Merge(&p, rt)
	return p
}
