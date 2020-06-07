package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Jagrmi-C/gostarted/project/db"
	"github.com/Jagrmi-C/gostarted/project/helpers"
	"github.com/Jagrmi-C/gostarted/project/router"
	"github.com/spf13/viper"
)

func readConfigFile() (error) {
	viper.SetConfigType("json")
	// TODO add path throuth library
	viper.SetConfigFile("project/config/myJSONConfig.json")
	fmt.Printf("Using config: %s\n", viper.ConfigFileUsed())
	err := viper.ReadInConfig()
	return err
}

type handler struct {
}

// type App struct {
// 	Router *mux.Router
// 	DB     *pgx.Conn
// 	Logger *log.Logger
// }

func (h *handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	res, _ := w.Write([]byte("Hello world!\n"))
	fmt.Println(res)
}

func main()  {
	err := readConfigFile()
	if err != nil {
		log.Fatal("!", err)
	}

	fmt.Println(viper.GetString("test"))

	PORT := ":8888"
	arguments := os.Args
	if len(arguments) == 1 {
		log.Println("Using default port number: ", PORT)
	} else {
		PORT = ":" + arguments[1]
	}

	app := helpers.App{}

	db.CheckDb()

	router := router.Router()
	app.Router = router

	fmt.Println("Listening to port number", PORT)
	if err := http.ListenAndServe(PORT, app.Router); err != nil {
		log.Fatal(err)
	}
}
