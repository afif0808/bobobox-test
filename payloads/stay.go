package payloads

type CreateStayPayload struct {
	RoomID    int64  `json:"room_id" merge:"room_id"`
	GuestName string `json:"guest_name" merge:"guest_name"`
}

type StayPayload struct {
	ID int64 `json:"id"`
}
