package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/user"

	lr "github.com/sirupsen/logrus"

	"github.com/Jagrmi-C/gostarted/project/db"
	"github.com/Jagrmi-C/gostarted/project/helpers"
	"github.com/Jagrmi-C/gostarted/project/router"
	"github.com/spf13/viper"

	"github.com/Jagrmi-C/gostarted/project/logger"
)

func init()  {
	logger.LoggerInitialization()
}

func startJobServer() {
	user, err := user.Current()
    if err != nil {
		lr.Error(err)
    }
	lr.WithFields(lr.Fields{
		"server":	viper.GetString("server"),
		"DBHost":	os.Getenv("DATABASE_URL"),
		"user":		user.Uid,
		"package":	"main",
	}).Info("Server start yours job")
}

func readConfigFile() (error) {
	viper.SetConfigType("json")
	// TODO add path throuth library
	viper.SetConfigFile("project/config/myJSONConfig.json")
	fmt.Printf("Using config: %s\n", viper.ConfigFileUsed())
	err := viper.ReadInConfig()
	return err
}

func main()  {
	err := readConfigFile()
	if err != nil {
		lr.Error(err)
	}

	startJobServer()

	PORT := ":8888"
	arguments := os.Args
	if len(arguments) == 1 {
		log.Println("Using default port number: ", PORT)
	} else {
		PORT = ":" + arguments[1]
	}

	app := helpers.App{}

	db.CreateConnection()

	router := router.Router()
	app.Router = router

	fmt.Println("Listening to port number", PORT)
	if err := http.ListenAndServe(PORT, app.Router); err != nil {
		log.Fatal(err)
	}
}
