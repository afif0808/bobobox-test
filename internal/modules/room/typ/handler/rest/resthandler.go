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
	rtrh := RoomTypeRestHandler{uc: uc}
	return &rtrh
}

func (rtrh *RoomTypeRestHandler) MountRoutes(e *echo.Echo) {
	e.POST("/room/type/", rtrh.CreateRoomType)
	e.PUT("/room/type/:id", rtrh.UpdateRoomType)
	e.GET("/room/type/", rtrh.GetRoomTypeList)
	e.GET("/room/type/:id", rtrh.GetRoomType)
	e.DELETE("/room/type/:id", rtrh.DeleteRoomType)

}

func (rtrh *RoomTypeRestHandler) CreateRoomType(c echo.Context) error {
	ctx := c.Request().Context()
	var payload payloads.CreateRoomTypePayload
	err := c.Bind(&payload)
	if err != nil {
		return wrapper.NewHTTPResponse(http.StatusBadRequest, "", nil, err).JSON(c.Response())
	}

	rt, err := rtrh.uc.CreateRoomType(ctx, payload)
	if err != nil {
		return wrapper.NewHTTPResponse(
			customerrors.ErrorHTTPCode(err), "", nil, err,
		).JSON(c.Response())
	}

	return wrapper.NewHTTPResponse(http.StatusCreated, "", rt.ToPayload(), nil).JSON(c.Response())
}

func (rtrh *RoomTypeRestHandler) GetRoomTypeList(c echo.Context) error {
	ctx := c.Request().Context()
	rts, err := rtrh.uc.GetRoomTypeList(ctx)
	if err != nil {
		return wrapper.NewHTTPResponse(
			customerrors.ErrorHTTPCode(err), "", nil, err,
		).JSON(c.Response())
	}
	results := make([]payloads.RoomTypePayload, len(rts))
	for i := range results {
		results[i] = rts[i].ToPayload()
	}

	return wrapper.NewHTTPResponse(http.StatusOK, "", results, nil).JSON(c.Response())
}

func (rtrh *RoomTypeRestHandler) UpdateRoomType(c echo.Context) error {
	ctx := c.Request().Context()
	var payload payloads.UpdateRoomTypePayload
	err := c.Bind(&payload)
	if err != nil {
		return wrapper.NewHTTPResponse(http.StatusBadRequest, "", nil, err).JSON(c.Response())
	}

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	rt, err := rtrh.uc.UpdateRoomType(ctx, payload, id)
	rt.ID = id

	if err != nil {
		log.Println(err)
		return wrapper.NewHTTPResponse(
			customerrors.ErrorHTTPCode(err), "", nil, err,
		).JSON(c.Response())
	}

	return wrapper.NewHTTPResponse(http.StatusOK, "", rt.ToPayload(), nil).JSON(c.Response())
}

func (rtrh *RoomTypeRestHandler) DeleteRoomType(c echo.Context) error {
	ctx := c.Request().Context()
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	err := rtrh.uc.DeleteRoomType(ctx, id)
	if err != nil {
		return wrapper.NewHTTPResponse(
			customerrors.ErrorHTTPCode(err), "", nil, err,
		).JSON(c.Response())
	}
	return wrapper.NewHTTPResponse(http.StatusOK, "", nil, nil).JSON(c.Response())
}

func (rtrh *RoomTypeRestHandler) GetRoomType(c echo.Context) error {
	ctx := c.Request().Context()
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	rt, err := rtrh.uc.GetRoomType(ctx, id)
	if err != nil {
		return wrapper.NewHTTPResponse(
			customerrors.ErrorHTTPCode(err), "", nil, err,
		).JSON(c.Response())
	}
	return wrapper.NewHTTPResponse(http.StatusOK, "", rt.ToPayload(), nil).JSON(c.Response())
}
