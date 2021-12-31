package price

import (
	"github.com/afif0808/bobobox_test/internal/modules/room/price/handler/rest"
	roompricerepo "github.com/afif0808/bobobox_test/internal/modules/room/price/repository"
	roomtyperepo "github.com/afif0808/bobobox_test/internal/modules/room/typ/repository"

	"github.com/afif0808/bobobox_test/internal/modules/room/price/usecase"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func InjectRoomPriceModule(e *echo.Echo, readDB, writeDB *sqlx.DB) {
	rpRepo := roompricerepo.NewRoomPriceSQLRepo(readDB, writeDB)
	rtRepo := roomtyperepo.NewRoomTypeSQLRepo(readDB, writeDB)

	uc := usecase.NewRoomPriceUsecase(rpRepo, rtRepo)
	hrh := rest.NewRoomPriceRestHandler(uc)
	hrh.MountRoutes(e)
}
