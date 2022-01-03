package main

import (
	"fmt"
	"log"
	"os"

	"github.com/afif0808/bobobox_test/internal/modules/hotel"
	"github.com/afif0808/bobobox_test/internal/modules/promo"
	"github.com/afif0808/bobobox_test/internal/modules/reservation"
	roomprice "github.com/afif0808/bobobox_test/internal/modules/room/price"
	"github.com/afif0808/bobobox_test/internal/modules/room/room"
	roomtype "github.com/afif0808/bobobox_test/internal/modules/room/typ"
	"github.com/afif0808/bobobox_test/internal/modules/stay"

	"github.com/afif0808/bobobox_test/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		return
	}
	e := echo.New()

	readDB, writeDB := initMySQLDB()
	migrateDatabase()

	hotel.InjectHotelModule(e, readDB, writeDB)
	roomtype.InjectRoomTypeModule(e, readDB, writeDB)
	roomprice.InjectRoomPriceModule(e, readDB, writeDB)
	room.InjectRoomModule(e, readDB, writeDB)
	reservation.InjectReservationModule(e, readDB, writeDB)
	promo.InjectPromoModule(e, readDB, writeDB)
	stay.InjectStayModule(e, readDB, writeDB)

	e.Start(":8080")
}
func initMySQLDB() (readDB, writeDB *sqlx.DB) {
	host := os.Getenv("MYSQL_HOST")
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	db := os.Getenv("MYSQL_DATABASE")
	port := os.Getenv("MYSQL_PORT")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&multiStatements=true", user, password, host, port, db)
	readDB, err := sqlx.Open("mysql", dsn)
	if err != nil {
		log.Panic(err)
	}
	err = readDB.Ping()
	if err != nil {
		log.Panic(err)
	}

	writeDB, err = sqlx.Open("mysql", dsn)
	if err != nil {
		log.Panic(err)
	}
	err = readDB.Ping()
	if err != nil {
		log.Panic(err)
	}

	return
}

func migrateDatabase() {
	host := os.Getenv("MYSQL_HOST")
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	db := os.Getenv("MYSQL_DATABASE")
	port := os.Getenv("MYSQL_PORT")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&multiStatements=true", user, password, host, port, db)
	conn, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Panic(err)
	}

	err = conn.AutoMigrate(
		&models.Hotel{},
		&models.RoomType{}, &models.RoomPrice{}, &models.Room{},
		&models.Reservation{}, &models.Stay{}, &models.StayDate{},
		&models.Promo{}, &models.PromoUse{},
	)
	if err != nil {
		log.Panic(err)
	}

}
