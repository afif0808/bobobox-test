package reservation

import (
	"github.com/afif0808/bobobox_test/internal/modules/reservation/handler/rest"
	"github.com/afif0808/bobobox_test/internal/modules/reservation/repository"
	"github.com/afif0808/bobobox_test/internal/modules/reservation/usecase"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func InjectReservationModule(e *echo.Echo, readDB, writeDB *sqlx.DB) {
	repo := repository.NewReservationSQLRepo(readDB, writeDB)
	uc := usecase.NewReservationUsecase(repo)
	hrh := rest.NewReservationRestHandler(uc)
	hrh.MountRoutes(e)
}
