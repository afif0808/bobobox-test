package stay

import (
	"github.com/afif0808/bobobox_test/internal/modules/stay/handler/rest"
	"github.com/afif0808/bobobox_test/internal/modules/stay/repository"
	"github.com/afif0808/bobobox_test/internal/modules/stay/usecase"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func InjectStayModule(e *echo.Echo, readDB, writeDB *sqlx.DB) {
	repo := repository.NewStaySQLRepo(readDB, writeDB)
	uc := usecase.NewStayUsecase(repo)
	hrh := rest.NewStayRestHandler(uc)
	hrh.MountRoutes(e)
}
