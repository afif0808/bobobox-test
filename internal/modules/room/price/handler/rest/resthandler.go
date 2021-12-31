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
	CreateRoomPrices(ctx context.Context, payload payloads.CreateRoomPricePayload) error
	UpdateRoomPrice(ctx context.Context, payload payloads.UpdateRoomPricePayload, id int64) error
	DeleteRoomPrice(ctx context.Context, id int64) error
	GetRoomTypePriceList(ctx context.Context, typeID int64) ([]models.RoomPrice, error)
	GetRoomPrice(ctx context.Context, id int64) (models.RoomPrice, error)
}

type RoomPriceRestHandler struct {
	uc Usecase
}

func NewRoomPriceRestHandler(uc Usecase) *RoomPriceRestHandler {
	rprh := RoomPriceRestHandler{uc: uc}
	return &rprh
}

func (rprh *RoomPriceRestHandler) MountRoutes(e *echo.Echo) {
	e.POST("/room/price/", rprh.CreateRoomPrices)
	e.PUT("/room/price/:id", rprh.UpdateRoomPrice)
	e.GET("/room/type/:id/price/", rprh.GetRoomTypePriceList)
	e.GET("/room/price/:id", rprh.GetRoomPrice)
	e.DELETE("/room/price/:id", rprh.DeleteRoomPrice)
}

func (rprh *RoomPriceRestHandler) CreateRoomPrices(c echo.Context) error {
	ctx := c.Request().Context()
	var payload payloads.CreateRoomPricePayload
	err := c.Bind(&payload)
	if err != nil {
		return wrapper.NewHTTPResponse(http.StatusBadRequest, "", nil, err).JSON(c.Response())
	}
	err = rprh.uc.CreateRoomPrices(ctx, payload)
	if err != nil {
		log.Println(err)
		return wrapper.NewHTTPResponse(
			customerrors.ErrorHTTPCode(err), "", nil, err,
		).JSON(c.Response())
	}
	return wrapper.NewHTTPResponse(http.StatusCreated, "", nil, nil).JSON(c.Response())
}

func (rprh *RoomPriceRestHandler) GetRoomTypePriceList(c echo.Context) error {
	ctx := c.Request().Context()
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	RoomPrices, err := rprh.uc.GetRoomTypePriceList(ctx, id)
	if err != nil {
		return wrapper.NewHTTPResponse(
			customerrors.ErrorHTTPCode(err), "", nil, err,
		).JSON(c.Response())
	}
	results := make([]payloads.RoomPricePayload, len(RoomPrices))
	for i := range results {
		results[i] = RoomPrices[i].ToPayload()
	}

	return wrapper.NewHTTPResponse(http.StatusOK, "", results, nil).JSON(c.Response())
}

func (rprh *RoomPriceRestHandler) UpdateRoomPrice(c echo.Context) error {
	ctx := c.Request().Context()
	var payload payloads.UpdateRoomPricePayload
	err := c.Bind(&payload)
	if err != nil {
		return wrapper.NewHTTPResponse(http.StatusBadRequest, "", nil, err).JSON(c.Response())
	}

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	err = rprh.uc.UpdateRoomPrice(ctx, payload, id)

	if err != nil {
		log.Println(err)
		return wrapper.NewHTTPResponse(
			customerrors.ErrorHTTPCode(err), "", nil, err,
		).JSON(c.Response())
	}

	return wrapper.NewHTTPResponse(http.StatusOK, "", nil, nil).JSON(c.Response())
}

func (rprh *RoomPriceRestHandler) DeleteRoomPrice(c echo.Context) error {
	ctx := c.Request().Context()
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	err := rprh.uc.DeleteRoomPrice(ctx, id)
	if err != nil {
		return wrapper.NewHTTPResponse(
			customerrors.ErrorHTTPCode(err), "", nil, err,
		).JSON(c.Response())
	}
	return wrapper.NewHTTPResponse(http.StatusOK, "", nil, nil).JSON(c.Response())
}

func (rprh *RoomPriceRestHandler) GetRoomPrice(c echo.Context) error {
	ctx := c.Request().Context()
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	rp, err := rprh.uc.GetRoomPrice(ctx, id)
	if err != nil {
		return wrapper.NewHTTPResponse(
			customerrors.ErrorHTTPCode(err), "", nil, err,
		).JSON(c.Response())
	}
	return wrapper.NewHTTPResponse(http.StatusOK, "", rp.ToPayload(), nil).JSON(c.Response())
}
