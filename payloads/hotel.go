package payloads

type CreateHotelPayload struct {
	Name    string `json:"name" merge:"name"`
	Address string `json:"address" merge:"address"`
}

type UpdateHotelPayload struct {
	Name    string `json:"name" merge:"name"`
	Address string `json:"address" merge:"address"`
}

type HotelPayload struct {
	ID      int64  `json:"id" merge:"id"`
	Name    string `json:"name" merge:"name"`
	Address string `json:"address" merge:"address"`
}
