package helpers

import (
	"log"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"
)

type App struct {
	Router *mux.Router
	DB     *pgx.Conn
	Logger *log.Logger
}
