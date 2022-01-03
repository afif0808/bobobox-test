package payloads

type CreateRoomTypePayload struct {
	Name         string  `json:"name" merge:"name" valid:"required"`
	DefaultPrice float64 `json:"default_price" merge:"default_price" valid:"required"`
}

type UpdateRoomTypePayload struct {
	Name         string  `json:"name" merge:"name" valid:"required"`
	DefaultPrice float64 `json:"default_price" merge:"default_price" valid:"required"`
}

type RoomTypePayload struct {
	ID           int64   `json:"id" merge:"id"`
	Name         string  `json:"name" merge:"name"`
	DefaultPrice float64 `json:"default_price" merge:"default_price"`
}
