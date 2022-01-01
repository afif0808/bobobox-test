package payloads

type CreateReservationPayload struct {
	CustomerName string              `json:"customer_name" merge:"customer_name" valid:"required"`
	Stays        []CreateStayPayload `json:"stays" valid:"required"`
	HotelID      int64               `json:"hotel_id" merge:"hotel_id" valid:"required"`
	CheckInDate  string              `json:"check_in_date" merge:"check_in_date" valid:"required"`
	CheckOutDate string              `json:"check_out_date" merge:"check_out_date" valid:"required"`
}

type ReservationPayload struct {
	ID              int64  `json:"id" merge:"id"`
	CheckInDate     string `json:"check_in_date" merge:"check_in_date"`
	CheckOutDate    string `json:"check_out_date" merge:"check_out_date"`
	BookedRoomCount int    `json:"booked_room_count" merge:"booked_room_count"`
	HotelID         int64  `json:"hotel_id" merge:"hotel_id"`
	CustomerName    string `json:"customer_name" merge:"customer_name"`
}
