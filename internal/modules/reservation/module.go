package reservation

import (
	promorepo "github.com/afif0808/bobobox_test/internal/modules/promo/repository"
	"github.com/afif0808/bobobox_test/internal/modules/reservation/handler/rest"
	reservationrepo "github.com/afif0808/bobobox_test/internal/modules/reservation/repository"
	"github.com/afif0808/bobobox_test/internal/modules/reservation/usecase"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func InjectReservationModule(e *echo.Echo, readDB, writeDB *sqlx.DB) {
	rRepo := reservationrepo.NewReservationSQLRepo(readDB, writeDB)
	prRepo := promorepo.NewPromoSQLRepo(readDB, writeDB)
	uc := usecase.NewReservationUsecase(rRepo, prRepo)
	hrh := rest.NewReservationRestHandler(uc)
	hrh.MountRoutes(e)
}
