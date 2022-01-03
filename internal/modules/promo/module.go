package promo

import (
	"github.com/afif0808/bobobox_test/internal/modules/promo/handler/rest"
	"github.com/afif0808/bobobox_test/internal/modules/promo/repository"
	"github.com/afif0808/bobobox_test/internal/modules/promo/usecase"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func InjectPromoModule(e *echo.Echo, readDB, writeDB *sqlx.DB) {
	repo := repository.NewPromoSQLRepo(readDB, writeDB)
	uc := usecase.NewPromoUsecase(repo)
	hrh := rest.NewPromoRestHandler(uc)
	hrh.MountRoutes(e)
}
