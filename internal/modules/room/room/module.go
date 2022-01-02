package room

import (
	hotelrepo "github.com/afif0808/bobobox_test/internal/modules/hotel/repository"
	roompricerepo "github.com/afif0808/bobobox_test/internal/modules/room/price/repository"
	"github.com/afif0808/bobobox_test/internal/modules/room/room/handler/rest"
	roomrepo "github.com/afif0808/bobobox_test/internal/modules/room/room/repository"
	"github.com/afif0808/bobobox_test/internal/modules/room/room/usecase"
	roomtyperepo "github.com/afif0808/bobobox_test/internal/modules/room/typ/repository"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func InjectRoomModule(e *echo.Echo, readDB, writeDB *sqlx.DB) {
	rRepo := roomrepo.NewRoomSQLRepo(readDB, writeDB)
	rtRepo := roomtyperepo.NewRoomTypeSQLRepo(readDB, writeDB)
	rpRepo := roompricerepo.NewRoomPriceSQLRepo(readDB, writeDB)
	hRepo := hotelrepo.NewHotelSQLRepo(readDB, writeDB)
	uc := usecase.NewRoomUsecase(rRepo, rtRepo, hRepo, rpRepo)
	hrh := rest.NewRoomRestHandler(uc)
	hrh.MountRoutes(e)
}
