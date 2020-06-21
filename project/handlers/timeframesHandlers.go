package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	lr "github.com/sirupsen/logrus"

	"github.com/Jagrmi-C/gostarted/project/db"
	"github.com/Jagrmi-C/gostarted/project/logger"
	"github.com/Jagrmi-C/gostarted/project/models"
)

func init()  {
	logger.LoggerInitialization()
}


func GetTimeframeHandler(w http.ResponseWriter, req *http.Request) {
	uuid := mux.Vars(req)["uuid"]

	var timeframe models.TimeFrame
	err := db.GetTimeFrame(uuid, &timeframe)
	if err != nil {
		lr.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	lr.Info("Get timeframe from DB:", timeframe.UUID)

	err = json.NewEncoder(w).Encode(timeframe)
	if err != nil {
		lr.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func CreateTimeframeHandler(w http.ResponseWriter, req *http.Request) {
	var timeframe models.TimeFrame
	err := json.NewDecoder(req.Body).Decode(&timeframe)
    if err != nil {
		lr.Error(err)
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
	}

	err = db.CreateTimeFrame(&timeframe)

	if err != nil {
		lr.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	lr.Info("Create timeframe:", timeframe)

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(timeframe)

	if err != nil {
		lr.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func DeleteTimeframeHandler(w http.ResponseWriter, req *http.Request) {
	uuid := mux.Vars(req)["uuid"]

	err := db.DeleteTimeFrame(uuid)
	if err != nil {
		lr.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	lr.Info("Delete timeframe from DB with uuid:", uuid)

	if err != nil {
		lr.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
