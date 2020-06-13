package helpers

import (
	"log"
	"time"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"
)

type App struct {
	Router *mux.Router
	DB     *pgx.Conn
	Logger *log.Logger
}

func GetCurrentLocalTime() time.Time {
	//init the loc
	loc, _ := time.LoadLocation("Europe/Minsk")

	//set timezone,  
	now := time.Now().In(loc)
	return now
}
