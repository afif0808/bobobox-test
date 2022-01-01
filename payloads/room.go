package payloads

type CreateRoomPayload struct {
	Number     string `json:"number" merge:"number"`
	HotelID    int64  `json:"hotel_id" merge:"hotel_id"`
	RoomTypeID int64  `json:"type_id" merge:"room_type_id"`
}

type UpdateRoomPayload struct {
	Number     string `json:"number" merge:"number"`
	RoomTypeID int64  `json:"type_id" merge:"room_type_id"`
}

type RoomPayload struct {
	ID          int64           `json:"id"  merge:"id"`
	Number      string          `json:"number"  merge:"number"`
	IsInService bool            `json:"is_in_service" merge:"is_in_service"`
	RoomType    RoomTypePayload `json:"type"`
	Hotel       HotelPayload    `json:"hotel"`
}
