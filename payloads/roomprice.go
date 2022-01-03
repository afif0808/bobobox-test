package payloads

type CreateRoomPricePayload struct {
	Date       string  `json:"date" merge:"date" valid:"required"`
	UntilDate  string  `json:"until_date"`
	Price      float64 `json:"price" merge:"price" valid:"required"`
	RoomTypeID int64   `json:"type_id" merge:"room_type_id" valid:"room_type_id"`
}
type UpdateRoomPricePayload struct {
	Price float64 `json:"price" merge:"price" valid:"required"`
}

type RoomPricePayload struct {
	ID         int64   `json:"id"  merge:"id"`
	Date       string  `json:"date"  merge:"date"`
	Price      float64 `json:"price"  merge:"price"`
	RoomTypeID int64   `json:"room_type_id"  merge:"room_type_id"`
}
