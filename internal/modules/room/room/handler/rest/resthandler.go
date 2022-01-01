package rest

import (
	"context"
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
	GetRoomList(tx context.Context) ([]models.Room, error)
	GetRoom(ctx context.Context, id int64) (models.Room, error)
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
	e.GET("/room/:id", rrh.GetRoom)
	e.DELETE("/room/:id", rrh.DeleteRoom)
}

func (rrh *RoomRestHandler) CreateRoom(c echo.Context) error {
	ctx := c.Request().Context()
	var payload payloads.CreateRoomPayload
	err := c.Bind(&payload)
	if err != nil {
		return wrapper.NewHTTPResponse(http.StatusBadRequest, "", nil, err).JSON(c.Response())
	}
	r, err := rrh.uc.CreateRoom(ctx, payload)
	if err != nil {
		log.Println(err)
		return wrapper.NewHTTPResponse(
			customerrors.ErrorHTTPCode(err), "", nil, err,
		).JSON(c.Response())
	}
	return wrapper.NewHTTPResponse(http.StatusCreated, "", r.ToPayload(), nil).JSON(c.Response())
}

func (rrh *RoomRestHandler) GetRoomList(c echo.Context) error {
	ctx := c.Request().Context()
	rs, err := rrh.uc.GetRoomList(ctx)
	if err != nil {
		log.Println(err)
		return wrapper.NewHTTPResponse(
			customerrors.ErrorHTTPCode(err), "", nil, err,
		).JSON(c.Response())
	}
	results := make([]payloads.RoomPayload, len(rs))
	for i := range results {
		results[i] = rs[i].ToPayload()
	}

	return wrapper.NewHTTPResponse(http.StatusOK, "", results, nil).JSON(c.Response())
}

func (rrh *RoomRestHandler) UpdateRoom(c echo.Context) error {
	ctx := c.Request().Context()
	var payload payloads.UpdateRoomPayload
	err := c.Bind(&payload)
	if err != nil {
		return wrapper.NewHTTPResponse(http.StatusBadRequest, "", nil, err).JSON(c.Response())
	}

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	err = rrh.uc.UpdateRoom(ctx, payload, id)

	if err != nil {
		log.Println(err)
		return wrapper.NewHTTPResponse(
			customerrors.ErrorHTTPCode(err), "", nil, err,
		).JSON(c.Response())
	}

	return wrapper.NewHTTPResponse(http.StatusOK, "", nil, nil).JSON(c.Response())
}

func (rrh *RoomRestHandler) DeleteRoom(c echo.Context) error {
	ctx := c.Request().Context()
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	err := rrh.uc.DeleteRoom(ctx, id)
	if err != nil {
		return wrapper.NewHTTPResponse(
			customerrors.ErrorHTTPCode(err), "", nil, err,
		).JSON(c.Response())
	}
	return wrapper.NewHTTPResponse(http.StatusOK, "", nil, nil).JSON(c.Response())
}

func (rrh *RoomRestHandler) GetRoom(c echo.Context) error {
	ctx := c.Request().Context()
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	r, err := rrh.uc.GetRoom(ctx, id)
	if err != nil {
		return wrapper.NewHTTPResponse(
			customerrors.ErrorHTTPCode(err), "", nil, err,
		).JSON(c.Response())
	}
	return wrapper.NewHTTPResponse(http.StatusOK, "", r.ToPayload(), nil).JSON(c.Response())
}
