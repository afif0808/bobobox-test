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
	CreateRoomType(ctx context.Context, payload payloads.CreateRoomTypePayload) (models.RoomType, error)
	UpdateRoomType(ctx context.Context, payload payloads.UpdateRoomTypePayload, id int64) (models.RoomType, error)
	DeleteRoomType(ctx context.Context, id int64) error
	GetRoomTypeList(ctx context.Context) ([]models.RoomType, error)
	GetRoomType(ctx context.Context, id int64) (models.RoomType, error)
}

type RoomTypeRestHandler struct {
	uc Usecase
}

func NewRoomTypeRestHandler(uc Usecase) *RoomTypeRestHandler {
	hrh := RoomTypeRestHandler{uc: uc}
	return &hrh
}

func (hrh *RoomTypeRestHandler) MountRoutes(e *echo.Echo) {
	e.POST("/room/type/", hrh.CreateRoomType)
	e.PUT("/room/type/:id", hrh.UpdateRoomType)
	e.GET("/room/type/", hrh.GetRoomTypeList)
	e.GET("/room/type/:id", hrh.GetRoomType)
	e.DELETE("/room/type/:id", hrh.DeleteRoomType)

}

func (hrh *RoomTypeRestHandler) CreateRoomType(c echo.Context) error {
	ctx := c.Request().Context()
	var payload payloads.CreateRoomTypePayload
	err := c.Bind(&payload)
	if err != nil {
		return wrapper.NewHTTPResponse(http.StatusBadRequest, "", nil, err).JSON(c.Response())
	}

	h, err := hrh.uc.CreateRoomType(ctx, payload)
	if err != nil {
		return wrapper.NewHTTPResponse(
			customerrors.ErrorHTTPCode(err), "", nil, err,
		).JSON(c.Response())
	}

	return wrapper.NewHTTPResponse(http.StatusCreated, "", h.ToPayload(), nil).JSON(c.Response())
}

func (hrh *RoomTypeRestHandler) GetRoomTypeList(c echo.Context) error {
	ctx := c.Request().Context()
	RoomTypes, err := hrh.uc.GetRoomTypeList(ctx)
	if err != nil {
		return wrapper.NewHTTPResponse(
			customerrors.ErrorHTTPCode(err), "", nil, err,
		).JSON(c.Response())
	}
	results := make([]payloads.RoomTypePayload, len(RoomTypes))
	for i := range results {
		results[i] = RoomTypes[i].ToPayload()
	}

	return wrapper.NewHTTPResponse(http.StatusOK, "", results, nil).JSON(c.Response())
}

func (hrh *RoomTypeRestHandler) UpdateRoomType(c echo.Context) error {
	ctx := c.Request().Context()
	var payload payloads.UpdateRoomTypePayload
	err := c.Bind(&payload)
	if err != nil {
		return wrapper.NewHTTPResponse(http.StatusBadRequest, "", nil, err).JSON(c.Response())
	}

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	RoomType, err := hrh.uc.UpdateRoomType(ctx, payload, id)
	RoomType.ID = id

	if err != nil {
		log.Println(err)
		return wrapper.NewHTTPResponse(
			customerrors.ErrorHTTPCode(err), "", nil, err,
		).JSON(c.Response())
	}

	return wrapper.NewHTTPResponse(http.StatusOK, "", RoomType.ToPayload(), nil).JSON(c.Response())
}

func (hrh *RoomTypeRestHandler) DeleteRoomType(c echo.Context) error {
	ctx := c.Request().Context()
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	err := hrh.uc.DeleteRoomType(ctx, id)
	if err != nil {
		return wrapper.NewHTTPResponse(
			customerrors.ErrorHTTPCode(err), "", nil, err,
		).JSON(c.Response())
	}
	return wrapper.NewHTTPResponse(http.StatusOK, "", nil, nil).JSON(c.Response())
}

func (hrh *RoomTypeRestHandler) GetRoomType(c echo.Context) error {
	ctx := c.Request().Context()
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	h, err := hrh.uc.GetRoomType(ctx, id)
	if err != nil {
		return wrapper.NewHTTPResponse(
			customerrors.ErrorHTTPCode(err), "", nil, err,
		).JSON(c.Response())
	}
	return wrapper.NewHTTPResponse(http.StatusOK, "", h.ToPayload(), nil).JSON(c.Response())
}
