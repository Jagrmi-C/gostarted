package main

import (
	"fmt"
	"net/http"
	"os"
	"os/user"

	lr "github.com/sirupsen/logrus"
	"github.com/joho/godotenv"

	"github.com/Jagrmi-C/gostarted/project/logger"
	"github.com/Jagrmi-C/gostarted/project/router"
)

func init()  {
	logger.LoggerInitialization()
}

func startJobServer() {
	if os.Getenv("writeLog") == "file" {
		writeLogFile()
	}

	user, err := user.Current()
    if err != nil {
		lr.Error(err)
    }
	lr.WithFields(lr.Fields{
		"server":	os.Getenv("server"),
		"DBHost":	os.Getenv("DATABASE_URL"),
		"user":		user.Uid,
		"package":	"main",
	}).Info("Server start yours job")
}

func getLogNameFile() (filename string) {
	filename = os.Getenv("fileLogName")
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
	err := godotenv.Load()
	if err != nil {
		lr.Error("Error loading .env file")
	}

	startJobServer()

	var PORT string
	arguments := os.Args
	if len(arguments) == 1 {
		PORT = fmt.Sprintf(":%s", os.Getenv("port"))
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
