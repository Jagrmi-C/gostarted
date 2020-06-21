package main

import (
	"fmt"
	"net/http"
	"os"
	"os/user"

	lr "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/Jagrmi-C/gostarted/project/logger"
	"github.com/Jagrmi-C/gostarted/project/router"
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
	viper.SetConfigFile("project/config/config.json")
	lr.Info(fmt.Sprintf("Using config: %s\n", viper.ConfigFileUsed()))
	err := viper.ReadInConfig()
	if viper.GetString("writeLog") == "file" {
		writeLogFile()
	}
	return err
}

func getLogNameFile() (filename string) {
	filename = viper.GetString("fileLogName")
	if filename != "" {
		filename += ".log"
	} else {
		filename = "logrus.log"
	}
	return filename
}

func writeLogFile()  {
	filename := getLogNameFile()
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0666)
    if err == nil {
        lr.SetOutput(file)
        lr.SetLevel(lr.InfoLevel)
    } else {
        lr.Info("Failed to log to file, using default stderr")
    }
	
}

func main()  {
	err := readConfigFile()
	if err != nil {
		lr.Error(err)
	}

	startJobServer()

	var PORT string
	arguments := os.Args
	if len(arguments) == 1 {
		PORT = fmt.Sprintf(":%s", viper.GetString("port"))
		fmt.Println("Using default port number: ", PORT)
	} else {
		PORT = ":" + arguments[1]
	}

	router := router.Router()
	fmt.Println("Listening to port number", PORT)
	if err := http.ListenAndServe(PORT, router); err != nil {
		lr.Fatal(err)
	}
}
