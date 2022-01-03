package payloads

type CreateRoomPayload struct {
	Number     string `json:"number" merge:"number" valid:"required"`
	HotelID    int64  `json:"hotel_id" merge:"hotel_id" valid:"required"`
	RoomTypeID int64  `json:"type_id" merge:"room_type_id" valid:"required"`
}

type UpdateRoomPayload struct {
	Number      string `json:"number" merge:"number" valid:"required"`
	RoomTypeID  int64  `json:"type_id" merge:"room_type_id" valid:"required"`
	IsInService bool   `json:"is_in_service" merge:"is_in_service" valid:"required"`
}

type RoomPayload struct {
	ID          int64           `json:"id"  merge:"id"`
	Number      string          `json:"number"  merge:"number"`
	IsInService bool            `json:"is_in_service" merge:"is_in_service"`
	RoomType    RoomTypePayload `json:"type"`
	Hotel       HotelPayload    `json:"hotel"`
}

type AvaiableRoomSummaryPayload struct {
	Message            string                 `json:"message"`
	CheckInDate        string                 `json:"check_id_date" merge:"check_in_date"`
	CheckOutDate       string                 `json:"check_out_date" merge:"check_out_date"`
	RoomTypeID         int64                  `json:"room_type_id" merge:"room_type_id"`
	RequestedRoomCount int                    `json:"requested_room_count" merge:"room_count"`
	TotalPrice         float64                `json:"total_price" merge:"total_price"`
	Rooms              []AvailableRoomPayload `json:"rooms"`
}

type AvailableRoomPayload struct {
	ID        int64                       `json:"id" merge:"id"`
	Number    string                      `json:"number" merge:"number"`
	Prices    []AvailableRoomPricePayload `json:"price" merge:"price"`
	HotelID   int64                       `json:"hotel_id" merge:"hotel_id"`
	HotelName string                      `json:"hotel_name" merge:"hotel_name"`
}

type AvailableRoomPricePayload struct {
	Price float64 `json:"price" merge:"price"`
	Date  string  `json:"date" merge:"date"`
}

type AvailableRoomInquiryPayload struct {
	RoomCount    int    `json:"room_count" merge:"room_count"`
	RoomTypeID   int64  `json:"room_type_id" merge:"room_type_id"`
	CheckInDate  string `json:"check_id_date" merge:"check_in_date"`
	CheckOutDate string `json:"check_out_date" merge:"check_out_date"`
}
