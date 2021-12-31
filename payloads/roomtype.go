package payloads

type CreateRoomTypePayload struct {
	Name string `json:"name" merge:"name"`
}
type UpdateRoomTypePayload struct {
	Name string `json:"name" merge:"name"`
}

type RoomTypePayload struct {
	ID   int64  `json:"id" merge:"id"`
	Name string `json:"name" merge:"name"`
}
