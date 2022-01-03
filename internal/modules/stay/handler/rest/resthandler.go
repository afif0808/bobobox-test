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
	GetReservationStayList(ctx context.Context, id int64) ([]models.Stay, error)
}

type StayRestHandler struct {
	uc Usecase
}

func NewStayRestHandler(uc Usecase) *StayRestHandler {
	hrh := StayRestHandler{uc: uc}
	return &hrh
}

func (hrh *StayRestHandler) MountRoutes(e *echo.Echo) {
	e.GET("/reservation/:id/stay/", hrh.GetReservationStayList)
}

func (hrh *StayRestHandler) GetReservationStayList(c echo.Context) error {
	ctx := c.Request().Context()
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	stays, err := hrh.uc.GetReservationStayList(ctx, id)
	if err != nil {
		return wrapper.NewHTTPResponse(
			customerrors.ErrorHTTPCode(err), nil, err,
		).JSON(c.Response())
	}

	results := make([]payloads.StayPayload, len(stays))
	for i := range results {
		results[i] = stays[i].ToPayload()
	}

	return wrapper.NewHTTPResponse(http.StatusOK, results, nil).JSON(c.Response())
}
