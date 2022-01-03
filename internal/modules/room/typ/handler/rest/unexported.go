package rest

import (
	"github.com/afif0808/bobobox_test/payloads"
	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
)

func getCreateRoomTypePayload(c echo.Context) (payloads.CreateRoomTypePayload, error) {
	var p payloads.CreateRoomTypePayload
	err := c.Bind(&p)
	if err != nil {
		return p, err
	}
	_, err = govalidator.ValidateStruct(p)
	return p, err
}

func getUpdateRoomTypePayload(c echo.Context) (payloads.UpdateRoomTypePayload, error) {
	var p payloads.UpdateRoomTypePayload
	err := c.Bind(&p)
	if err != nil {
		return p, err
	}
	_, err = govalidator.ValidateStruct(p)
	return p, err
}
