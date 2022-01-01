package rest

import (
	"errors"

	"github.com/afif0808/bobobox_test/payloads"
	"github.com/afif0808/bobobox_test/pkg/customerrors"
	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
)

func getCreatePayload(c echo.Context) (payloads.CreateReservationPayload, error) {
	var p payloads.CreateReservationPayload
	err := c.Bind(&p)
	if err != nil {
		return p, err
	}
	ok, err := govalidator.ValidateStruct(p)
	if err != nil || !ok {
		return p, err
	}

	if p.CheckInDate > p.CheckOutDate {
		return p, customerrors.NewCustomError(
			"bad payload",
			errors.New("check in date cannot be later than check out date"),
			customerrors.ErrTypeBadRequest,
		)
	}

	return p, nil
}
