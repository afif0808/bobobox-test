package models

import (
	"strconv"
	"strings"
	"time"

	"github.com/afif0808/bobobox_test/payloads"
	"github.com/afif0808/bobobox_test/pkg/structs"
)

type Promo struct {
	ID            int64   `db:"id"  merge:"id" action:"create" gorm:"primaryKey;autoIncrement:false"`
	Type          string  `db:"type" action:"create,update" merge:"type"` // amount or percentage
	Value         float64 `db:"value" action:"create,update" merge:"value"`
	CheckInDays   string  `db:"check_in_days" action:"create,update" ` // day number separated by coma , eg : 0,1,2,3,etc
	BookingDays   string  `db:"booking_days" action:"create,update"`   // day number separated by coma , eg  : 0,1,2,3,etc
	BookingHour   string  `db:"booking_hour" action:"create,update" merge:"booking_hour"`
	Quota         int     `db:"quota" action:"create,update" merge:"quota"`
	DailyMaxQuota int     `db:"daily_max_quota" action:"create,update" merge:"daily_max_quota"`
	MinimumNight  int     `db:"minimum_night" action:"create,update" merge:"minimum_night"`
	MinimumRoom   int     `db:"minimum_room" action:"create,update" merge:"minimum_room"`
}

type PromoUse struct {
	ReservationID int64  `db:"reservation_id" action:"create"`
	PromoID       int64  `db:"promo_id" action:"create"`
	BookingDate   string `db:"booking_date" action:"create" gorm:"type:date"`
}

func (pr Promo) ToPayload() payloads.PromoPayload {
	var p payloads.PromoPayload
	structs.Merge(&p, pr)

	for _, v := range strings.Split(pr.BookingDays, ",") {
		i, _ := strconv.Atoi(v)
		p.BookingDays = append(p.BookingDays, time.Weekday(i))
	}

	for _, v := range strings.Split(pr.CheckInDays, ",") {
		i, _ := strconv.Atoi(v)
		p.CheckInDays = append(p.CheckInDays, time.Weekday(i))
	}
	return p
}
