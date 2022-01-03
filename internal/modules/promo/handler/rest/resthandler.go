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
	CreatePromo(ctx context.Context, payload payloads.CreatePromoPayload) (models.Promo, error)
	UpdatePromo(ctx context.Context, payload payloads.UpdatePromoPayload, id int64) (models.Promo, error)
	DeletePromo(ctx context.Context, id int64) error
	GetPromoList(ctx context.Context) ([]models.Promo, error)
	GetPromo(ctx context.Context, id int64) (models.Promo, error)
	CheckPromo(ctx context.Context, p payloads.CheckPromoPayload) (payloads.CheckPromoResultPayload, error)
}

type PromoRestHandler struct {
	uc Usecase
}

func NewPromoRestHandler(uc Usecase) *PromoRestHandler {
	prh := PromoRestHandler{uc: uc}
	return &prh
}

func (prh *PromoRestHandler) MountRoutes(e *echo.Echo) {
	e.POST("/promo/", prh.CreatePromo)
	e.PUT("/promo/:id", prh.UpdatePromo)
	e.GET("/promo/", prh.GetPromoList)
	e.GET("/promo/:id", prh.GetPromo)
	e.DELETE("/promo/:id", prh.DeletePromo)
	e.POST("/promo/check/", prh.CheckPromo)

}

func (prh *PromoRestHandler) CreatePromo(c echo.Context) error {
	ctx := c.Request().Context()
	payload, err := getCreatePromoPayload(c)
	if err != nil {
		return wrapper.NewHTTPResponse(http.StatusBadRequest, nil, err).JSON(c.Response())
	}

	pr, err := prh.uc.CreatePromo(ctx, payload)
	if err != nil {
		return wrapper.NewHTTPResponse(
			customerrors.ErrorHTTPCode(err), nil, err,
		).JSON(c.Response())
	}

	return wrapper.NewHTTPResponse(http.StatusCreated, pr.ToPayload(), nil).JSON(c.Response())
}

func (prh *PromoRestHandler) GetPromoList(c echo.Context) error {
	ctx := c.Request().Context()
	promos, err := prh.uc.GetPromoList(ctx)
	if err != nil {
		return wrapper.NewHTTPResponse(
			customerrors.ErrorHTTPCode(err), nil, err,
		).JSON(c.Response())
	}
	results := make([]payloads.PromoPayload, len(promos))
	for i := range results {
		results[i] = promos[i].ToPayload()
	}

	return wrapper.NewHTTPResponse(http.StatusOK, results, nil).JSON(c.Response())
}

func (prh *PromoRestHandler) UpdatePromo(c echo.Context) error {
	ctx := c.Request().Context()
	payload, err := getUpdatePromoPayload(c)
	if err != nil {
		return wrapper.NewHTTPResponse(http.StatusBadRequest, nil, err).JSON(c.Response())
	}

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	promo, err := prh.uc.UpdatePromo(ctx, payload, id)
	promo.ID = id

	if err != nil {
		return wrapper.NewHTTPResponse(
			customerrors.ErrorHTTPCode(err), nil, err,
		).JSON(c.Response())
	}

	return wrapper.NewHTTPResponse(http.StatusOK, promo.ToPayload(), nil).JSON(c.Response())
}

func (prh *PromoRestHandler) DeletePromo(c echo.Context) error {
	ctx := c.Request().Context()
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	err := prh.uc.DeletePromo(ctx, id)
	if err != nil {
		return wrapper.NewHTTPResponse(
			customerrors.ErrorHTTPCode(err), nil, err,
		).JSON(c.Response())
	}
	return wrapper.NewHTTPResponse(http.StatusOK, nil, nil).JSON(c.Response())
}

func (prh *PromoRestHandler) GetPromo(c echo.Context) error {
	ctx := c.Request().Context()
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	pr, err := prh.uc.GetPromo(ctx, id)
	if err != nil {
		return wrapper.NewHTTPResponse(
			customerrors.ErrorHTTPCode(err), nil, err,
		).JSON(c.Response())
	}
	return wrapper.NewHTTPResponse(http.StatusOK, pr.ToPayload(), nil).JSON(c.Response())
}

func (prh *PromoRestHandler) getCheckPromoPaylaod(c echo.Context) (payloads.CheckPromoPayload, error) {
	// NOTE : I decided not to have 'total_price' field
	var p payloads.CheckPromoPayload
	c.Bind(&p)

	return p, nil
}

func (prh *PromoRestHandler) CheckPromo(c echo.Context) error {
	ctx := c.Request().Context()
	p, err := prh.getCheckPromoPaylaod(c)
	if err != nil {
		return wrapper.NewHTTPResponse(http.StatusBadRequest, nil, err).JSON(c.Response())
	}
	result, err := prh.uc.CheckPromo(ctx, p)
	if err != nil {
		return wrapper.NewHTTPResponse(
			customerrors.ErrorHTTPCode(err), nil, err,
		).JSON(c.Response())
	}
	return wrapper.NewHTTPResponse(http.StatusOK, result, err).JSON(c.Response())
}
