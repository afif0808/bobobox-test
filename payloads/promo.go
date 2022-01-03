package payloads

import (
	"time"
)

type PromoPayload struct {
	ID            int64          `json:"id" merge:"id"`
	Type          string         `json:"type" merge:"type" ` // amount or percentage
	Value         float64        `json:"value" merge:"value" `
	CheckInDays   []time.Weekday `json:"check_in_days" `
	BookingDays   []time.Weekday `json:"booking_days" `
	BookingHour   string         `json:"booking_hour" merge:"booking_hour" ` // eg: 08.00-23.00
	Quota         int            `json:"quota" merge:"quota"`
	DailyMaxQuota int            `json:"daily_quota" merge:"daily_max_quota" `
	MinimumNight  int            `json:"minimum_night" merge:"minimum_night"`
	MinimumRoom   int            `json:"minimum_room" merge:"minimum_room"`
}

type CreatePromoPayload struct {
	Type          string         `json:"type" merge:"type" valid:"required"` // amount or percentage
	Value         float64        `json:"value" merge:"value" valid:"required"`
	CheckInDays   []time.Weekday `json:"check_in_days"  valid:"required"`
	BookingDays   []time.Weekday `json:"booking_days" valid:"required" `
	BookingHour   string         `json:"booking_hour" merge:"booking_hour" valid:"required"`
	Quota         int            `json:"quota" merge:"quota" valid:"required"`
	DailyMaxQuota int            `json:"daily_max_quota" merge:"daily_max_quota" valid:"required"`
	MinimumNight  int            `json:"minimum_night" merge:"minimum_night" valid:"required"`
	MinimumRoom   int            `json:"minimum_room" merge:"minimum_room" valid:"required"`
}

type UpdatePromoPayload struct {
	Type          string         `json:"type" merge:"type" valid:"required"` // amount or percentage
	Value         float64        `json:"value" merge:"value" valid:"required"`
	CheckInDays   []time.Weekday `json:"check_in_days"  valid:"required"`
	BookingDays   []time.Weekday `json:"booking_days"  valid:"required"`
	BookingHour   string         `json:"booking_hour" merge:"booking_hour" valid:"required"`
	Quota         int            `json:"quota" merge:"quota" valid:"required"`
	DailyMaxQuota int            `json:"daily_max_quota" merge:"daily_max_quota" valid:"required"`
	MinimumNight  int            `json:"minimum_night" merge:"minimum_night" valid:"required"`
	MinimumRoom   int            `json:"minimum_room" merge:"minimum_room" valid:"required"`
}

type CheckPromoPayload struct {
	PromoID int64            `json:"promo_id"`
	Rooms   []CheckPromoRoom `json:"rooms"`
	// TotalPrice float64 `json:"total_price"`
	// Decided not to use total price field in order to
	// apply single source of truth principle, so total price will be calculated in the backend
}

type CheckPromoRoom struct {
	CheckInDate  string  `json:"check_in_date"`
	CheckOutDate string  `json:"check_out_date"`
	Price        float64 `json:"price"`
}

type CheckPromoResultPayload struct {
	DefaultTotalPrice float64          `json:"default_total_price"`
	FinalPrice        float64          `json:"final_price"`
	TotalPromo        float64          `json:"total_promo"`
	Rooms             []CheckPromoRoom `json:"rooms"`
}
