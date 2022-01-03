package payloads

type CreateStayPayload struct {
	RoomID    int64  `json:"room_id" merge:"room_id"`
	GuestName string `json:"guest_name" merge:"guest_name"`
}

type StayPayload struct {
	ID            int64              `json:"id" merge:"id" `
	RoomID        int64              `json:"room_id" merge:"room_id"`
	GuestName     string             `json:"guest_name" merge:"guest_name"`
	ReservationID int64              `json:"reservation_id" merge:"reservation_id"`
	Reservation   ReservationPayload `json:"reservation" `
	Dates         []StayDatePayload  `json:"dates"`
}

type StayDatePayload struct {
	Date string `json:"date"`
}
