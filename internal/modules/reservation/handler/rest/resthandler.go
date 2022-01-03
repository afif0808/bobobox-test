package rest

import (
	"context"
	"net/http"
	"strconv"

	"github.com/afif0808/bobobox_test/models"
	"github.com/afif0808/bobobox_test/payloads"
	"github.com/afif0808/bobobox_test/pkg/customerrors"
	"github.com/afif0808/bobobox_test/pkg/wrapper"
	"github.com/labstack/echo/v4"
)

type Usecase interface {
	CreateReservation(ctx context.Context, payload payloads.CreateReservationPayload) (models.Reservation, error)
	DeleteReservation(ctx context.Context, id int64) error
	GetReservationList(ctx context.Context) ([]models.Reservation, error)
	GetReservation(ctx context.Context, id int64) (models.Reservation, error)
}

type ReservationRestHandler struct {
	uc Usecase
}

func NewReservationRestHandler(uc Usecase) *ReservationRestHandler {
	hrh := ReservationRestHandler{uc: uc}
	return &hrh
}

func (hrh *ReservationRestHandler) MountRoutes(e *echo.Echo) {
	e.POST("/reservation/", hrh.CreateReservation)
	e.GET("/reservation/", hrh.GetReservationList)
	e.GET("/reservation/:id", hrh.GetReservation)
	e.DELETE("/reservation/:id", hrh.DeleteReservation)
}

func (hrh *ReservationRestHandler) CreateReservation(c echo.Context) error {
	ctx := c.Request().Context()
	p, err := getCreatePayload(c)
	if err != nil {
		return wrapper.NewHTTPResponse(http.StatusBadRequest, nil, err).JSON(c.Response())
	}

	re, err := hrh.uc.CreateReservation(ctx, p)
	if err != nil {
		return wrapper.NewHTTPResponse(
			customerrors.ErrorHTTPCode(err), nil, err,
		).JSON(c.Response())
	}

	return wrapper.NewHTTPResponse(http.StatusCreated, re.ToPayload(), nil).JSON(c.Response())
}

func (hrh *ReservationRestHandler) GetReservationList(c echo.Context) error {
	ctx := c.Request().Context()
	reservations, err := hrh.uc.GetReservationList(ctx)
	if err != nil {
		return wrapper.NewHTTPResponse(
			customerrors.ErrorHTTPCode(err), nil, err,
		).JSON(c.Response())
	}
	results := make([]payloads.ReservationPayload, len(reservations))
	for i := range results {
		results[i] = reservations[i].ToPayload()
	}

	return wrapper.NewHTTPResponse(http.StatusOK, results, nil).JSON(c.Response())
}

func (hrh *ReservationRestHandler) DeleteReservation(c echo.Context) error {
	ctx := c.Request().Context()
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	err := hrh.uc.DeleteReservation(ctx, id)
	if err != nil {
		return wrapper.NewHTTPResponse(
			customerrors.ErrorHTTPCode(err), nil, err,
		).JSON(c.Response())
	}
	return wrapper.NewHTTPResponse(http.StatusOK, nil, nil).JSON(c.Response())
}

func (hrh *ReservationRestHandler) GetReservation(c echo.Context) error {
	ctx := c.Request().Context()
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	h, err := hrh.uc.GetReservation(ctx, id)
	if err != nil {
		return wrapper.NewHTTPResponse(
			customerrors.ErrorHTTPCode(err), nil, err,
		).JSON(c.Response())
	}
	return wrapper.NewHTTPResponse(http.StatusOK, h.ToPayload(), nil).JSON(c.Response())
}
