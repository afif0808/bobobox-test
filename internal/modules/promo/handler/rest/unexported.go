package rest

import (
	"github.com/afif0808/bobobox_test/payloads"
	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
)

func getCreatePromoPayload(c echo.Context) (payloads.CreatePromoPayload, error) {
	var p payloads.CreatePromoPayload
	err := c.Bind(&p)
	if err != nil {
		return p, err
	}
	_, err = govalidator.ValidateStruct(p)
	return p, err
}

func getUpdatePromoPayload(c echo.Context) (payloads.UpdatePromoPayload, error) {
	var p payloads.UpdatePromoPayload
	err := c.Bind(&p)
	if err != nil {
		return p, err
	}
	_, err = govalidator.ValidateStruct(p)
	return p, err
}
