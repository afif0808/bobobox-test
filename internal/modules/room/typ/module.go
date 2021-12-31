package typ

import (
	"github.com/afif0808/bobobox_test/internal/modules/room/typ/handler/rest"
	"github.com/afif0808/bobobox_test/internal/modules/room/typ/repository"
	"github.com/afif0808/bobobox_test/internal/modules/room/typ/usecase"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func InjectRoomTypeModule(e *echo.Echo, readDB, writeDB *sqlx.DB) {
	repo := repository.NewRoomTypeSQLRepo(readDB, writeDB)
	uc := usecase.NewRoomTypeUsecase(repo)
	hrh := rest.NewRoomTypeRestHandler(uc)
	hrh.MountRoutes(e)
}
