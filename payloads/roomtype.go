package payloads

type CreateRoomTypePayload struct {
	Name         string  `json:"name" merge:"name"`
	DefaultPrice float64 `json:"default_price" merge:"default_price"`
}

type UpdateRoomTypePayload struct {
	Name         string  `json:"name" merge:"name"`
	DefaultPrice float64 `json:"default_price" merge:"default_price"`
}

type RoomTypePayload struct {
	ID           int64   `json:"id" merge:"id"`
	Name         string  `json:"name" merge:"name"`
	DefaultPrice float64 `json:"default_price" merge:"default_price"`
}
