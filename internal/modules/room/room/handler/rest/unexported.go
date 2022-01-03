package rest

import (
	"github.com/afif0808/bobobox_test/payloads"
	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
)

func getCreateRoomPayload(c echo.Context) (payloads.CreateRoomPayload, error) {
	var p payloads.CreateRoomPayload
	err := c.Bind(&p)
	if err != nil {
		return p, err
	}
	_, err = govalidator.ValidateStruct(p)
	return p, err
}

func getUpdateRoomPayload(c echo.Context) (payloads.UpdateRoomPayload, error) {
	var p payloads.UpdateRoomPayload
	err := c.Bind(&p)
	if err != nil {
		return p, err
	}
	_, err = govalidator.ValidateStruct(p)
	return p, err
}
