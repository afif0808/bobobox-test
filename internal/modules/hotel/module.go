package hotel

import (
	"github.com/afif0808/bobobox_test/internal/modules/hotel/handler/rest"
	"github.com/afif0808/bobobox_test/internal/modules/hotel/repository"
	"github.com/afif0808/bobobox_test/internal/modules/hotel/usecase"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func InjectHotelModule(e *echo.Echo, readDB, writeDB *sqlx.DB) {
	repo := repository.NewHotelSQLRepo(readDB, writeDB)
	uc := usecase.NewHotelUsecase(repo)
	hrh := rest.NewHotelRestHandler(uc)
	hrh.MountRoutes(e)
}
