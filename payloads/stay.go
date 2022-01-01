package payloads

type CreateStayPayload struct {
	RoomID    int64  `json:"room_id" merge:"room_id"`
	GuestName string `json:"guest_name" merge:"guest_name"`
}
