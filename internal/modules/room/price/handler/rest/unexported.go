package rest

import (
	"time"

	"github.com/afif0808/bobobox_test/payloads"
	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
)

func getCreateRoomPricePayload(c echo.Context) (payloads.CreateRoomPricePayload, error) {
	var p payloads.CreateRoomPricePayload
	err := c.Bind(&p)
	if err != nil {
		return p, err
	}
	_, err = govalidator.ValidateStruct(p)
	if err != nil {
		return p, err
	}

	_, err = time.Parse("2006-01-02", p.Date)
	if p.UntilDate != "" {
		_, err = time.Parse("2006-01-02", p.UntilDate)
	}

	return p, err
}

func getUpdateRoomPricePayload(c echo.Context) (payloads.UpdateRoomPricePayload, error) {
	var p payloads.UpdateRoomPricePayload
	err := c.Bind(&p)
	if err != nil {
		return p, err
	}
	_, err = govalidator.ValidateStruct(p)
	return p, err
}
