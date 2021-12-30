package rest

import (
	"context"
	"net/http"
	"strconv"

	"github.com/afif0808/bobobox_test/models"
	"github.com/afif0808/bobobox_test/payloads"
	"github.com/labstack/echo/v4"
)

type Usecase interface {
	CreateHotel(ctx context.Context, payload payloads.CreateHotelPayload) (models.Hotel, error)
	UpdateHotel(ctx context.Context, payload payloads.UpdateHotelPayload, id int64) (models.Hotel, error)
	DeleteHotel(ctx context.Context, id int64) error
	GetHotelList(ctx context.Context) ([]models.Hotel, error)
}

type HotelRestHandler struct {
	uc Usecase
}

func NewHotelRestHandler(uc Usecase) *HotelRestHandler {
	hrh := HotelRestHandler{uc: uc}
	return &hrh
}

func (hrh *HotelRestHandler) MountRoutes(e *echo.Echo) {
	e.POST("/hotel/", hrh.CreateHotel)
	e.PUT("/hotel/:id", hrh.UpdateHotel)
	e.GET("/hotel/", hrh.GetHotelList)
	e.DELETE("/hotel/:id", hrh.DeleteHotel)

}

func (hrh *HotelRestHandler) CreateHotel(c echo.Context) error {
	ctx := c.Request().Context()
	var payload payloads.CreateHotelPayload
	err := c.Bind(&payload)
	if err != nil {
		// error handling
	}

	hotel, err := hrh.uc.CreateHotel(ctx, payload)
	if err != nil {
		// error handling
	}

	return c.JSON(http.StatusCreated, hotel.ToPayload())
}

func (hrh *HotelRestHandler) GetHotelList(c echo.Context) error {
	ctx := c.Request().Context()
	hotels, err := hrh.uc.GetHotelList(ctx)
	if err != nil {
		// error handling
	}
	results := make([]payloads.HotelPayload, len(hotels))
	for i := range results {
		results[i] = hotels[i].ToPayload()
	}

	return c.JSON(http.StatusOK, results)
}

func (hrh *HotelRestHandler) UpdateHotel(c echo.Context) error {
	ctx := c.Request().Context()
	var payload payloads.UpdateHotelPayload
	err := c.Bind(&payload)
	if err != nil {
		// error handling
	}

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	hotel, err := hrh.uc.UpdateHotel(ctx, payload, id)
	hotel.ID = id

	if err != nil {
		// error handling
	}

	return c.JSON(http.StatusCreated, hotel.ToPayload())
}

func (hrh *HotelRestHandler) DeleteHotel(c echo.Context) error {
	ctx := c.Request().Context()
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	err := hrh.uc.DeleteHotel(ctx, id)
	if err != nil {

	}
	return c.JSON(http.StatusOK, nil)
}
