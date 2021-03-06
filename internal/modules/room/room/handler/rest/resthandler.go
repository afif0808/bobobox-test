package rest

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/afif0808/bobobox_test/models"
	"github.com/afif0808/bobobox_test/payloads"
	"github.com/afif0808/bobobox_test/pkg/customerrors"
	"github.com/afif0808/bobobox_test/pkg/wrapper"
	"github.com/labstack/echo/v4"
)

type Usecase interface {
	CreateRoom(ctx context.Context, payload payloads.CreateRoomPayload) (models.Room, error)
	UpdateRoom(ctx context.Context, payload payloads.UpdateRoomPayload, id int64) error
	DeleteRoom(ctx context.Context, id int64) error
	GetRoomList(ctx context.Context) ([]models.Room, error)
	GetHotelRoomList(ctx context.Context, id int64) ([]models.Room, error)
	GetRoom(ctx context.Context, id int64) (models.Room, error)
	GetAvailableRooms(ctx context.Context, p payloads.AvailableRoomInquiryPayload) (payloads.AvaiableRoomSummaryPayload, error)
}

type RoomRestHandler struct {
	uc Usecase
}

func NewRoomRestHandler(uc Usecase) *RoomRestHandler {
	rrh := RoomRestHandler{uc: uc}
	return &rrh
}

func (rrh *RoomRestHandler) MountRoutes(e *echo.Echo) {
	e.POST("/room/", rrh.CreateRoom)
	e.PUT("/room/:id", rrh.UpdateRoom)
	e.GET("/room/", rrh.GetRoomList)
	e.GET("/hotel/:id/room/", rrh.GetHotelRoomList)
	e.GET("/room/:id", rrh.GetRoom)
	e.DELETE("/room/:id", rrh.DeleteRoom)
	e.GET("/room/available/", rrh.GetAvailableRooms)

}

func (rrh *RoomRestHandler) CreateRoom(c echo.Context) error {
	ctx := c.Request().Context()
	payload, err := getCreateRoomPayload(c)
	if err != nil {
		return wrapper.NewHTTPResponse(http.StatusBadRequest, nil, err).JSON(c.Response())
	}
	r, err := rrh.uc.CreateRoom(ctx, payload)
	if err != nil {
		log.Println(err)
		return wrapper.NewHTTPResponse(
			customerrors.ErrorHTTPCode(err), nil, err,
		).JSON(c.Response())
	}
	return wrapper.NewHTTPResponse(http.StatusCreated, r.ToPayload(), nil).JSON(c.Response())
}

func (rrh *RoomRestHandler) GetRoomList(c echo.Context) error {
	ctx := c.Request().Context()
	rs, err := rrh.uc.GetRoomList(ctx)
	if err != nil {
		log.Println(err)
		return wrapper.NewHTTPResponse(
			customerrors.ErrorHTTPCode(err), nil, err,
		).JSON(c.Response())
	}
	results := make([]payloads.RoomPayload, len(rs))
	for i := range results {
		results[i] = rs[i].ToPayload()
	}

	return wrapper.NewHTTPResponse(http.StatusOK, results, nil).JSON(c.Response())
}

func (rrh *RoomRestHandler) GetHotelRoomList(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	ctx := c.Request().Context()
	rs, err := rrh.uc.GetHotelRoomList(ctx, id)
	if err != nil {
		log.Println(err)
		return wrapper.NewHTTPResponse(
			customerrors.ErrorHTTPCode(err), nil, err,
		).JSON(c.Response())
	}
	results := make([]payloads.RoomPayload, len(rs))
	for i := range results {
		results[i] = rs[i].ToPayload()
	}

	return wrapper.NewHTTPResponse(http.StatusOK, results, nil).JSON(c.Response())
}

func (rrh *RoomRestHandler) UpdateRoom(c echo.Context) error {
	ctx := c.Request().Context()
	payload, err := getUpdateRoomPayload(c)
	if err != nil {
		return wrapper.NewHTTPResponse(http.StatusBadRequest, nil, err).JSON(c.Response())
	}

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	err = rrh.uc.UpdateRoom(ctx, payload, id)

	if err != nil {
		log.Println(err)
		return wrapper.NewHTTPResponse(
			customerrors.ErrorHTTPCode(err), nil, err,
		).JSON(c.Response())
	}

	return wrapper.NewHTTPResponse(http.StatusOK, nil, nil).JSON(c.Response())
}

func (rrh *RoomRestHandler) DeleteRoom(c echo.Context) error {
	ctx := c.Request().Context()
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	err := rrh.uc.DeleteRoom(ctx, id)
	if err != nil {
		return wrapper.NewHTTPResponse(
			customerrors.ErrorHTTPCode(err), nil, err,
		).JSON(c.Response())
	}
	return wrapper.NewHTTPResponse(http.StatusOK, nil, nil).JSON(c.Response())
}

func (rrh *RoomRestHandler) GetRoom(c echo.Context) error {
	ctx := c.Request().Context()
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	r, err := rrh.uc.GetRoom(ctx, id)
	if err != nil {
		return wrapper.NewHTTPResponse(
			customerrors.ErrorHTTPCode(err), nil, err,
		).JSON(c.Response())
	}
	return wrapper.NewHTTPResponse(http.StatusOK, r.ToPayload(), nil).JSON(c.Response())
}

func (rrh *RoomRestHandler) getAvailableRoomInquiryPayload(c echo.Context) (payloads.AvailableRoomInquiryPayload, error) {
	var p payloads.AvailableRoomInquiryPayload
	p.CheckInDate = c.QueryParam("check_in_date")
	p.CheckOutDate = c.QueryParam("check_out_date")
	p.RoomCount, _ = strconv.Atoi(c.QueryParam("room_count"))
	p.RoomTypeID, _ = strconv.ParseInt(c.QueryParam("room_type_id"), 10, 64)
	if p.CheckInDate == "" || p.CheckOutDate == "" {
		return p, customerrors.NewCustomError("missing requirement", errors.New("check in date and chec out date is required"), customerrors.ErrTypeBadRequest)
	}

	if p.CheckInDate >= p.CheckOutDate {
		return p, customerrors.NewCustomError("bad request", errors.New("check in date cannot be same or later to check out date"), customerrors.ErrTypeBadRequest)
	}

	return p, nil
}

func (rrh *RoomRestHandler) GetAvailableRooms(c echo.Context) error {
	ctx := c.Request().Context()
	p, err := rrh.getAvailableRoomInquiryPayload(c)
	if err != nil {
		return wrapper.NewHTTPResponse(http.StatusBadRequest, nil, err).JSON(c.Response())
	}
	summary, err := rrh.uc.GetAvailableRooms(ctx, p)
	if err != nil {
		log.Println(err)
		return wrapper.NewHTTPResponse(
			customerrors.ErrorHTTPCode(err), nil, err,
		).JSON(c.Response())
	}

	return wrapper.NewHTTPResponse(http.StatusOK, summary, nil).JSON(c.Response())
}
