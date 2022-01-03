package rest

import (
	"github.com/afif0808/bobobox_test/payloads"
	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
)

func getCreateHotelPayload(c echo.Context) (payloads.CreateHotelPayload, error) {
	var p payloads.CreateHotelPayload
	err := c.Bind(&p)
	if err != nil {
		return p, err
	}
	_, err = govalidator.ValidateStruct(p)
	return p, err
}

func getUpdateHotelPayload(c echo.Context) (payloads.UpdateHotelPayload, error) {
	var p payloads.UpdateHotelPayload
	err := c.Bind(&p)
	if err != nil {
		return p, err
	}
	_, err = govalidator.ValidateStruct(p)
	return p, err
}
