package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	lr "github.com/sirupsen/logrus"

	"github.com/Jagrmi-C/gostarted/project/db"
	"github.com/Jagrmi-C/gostarted/project/helpers"
	"github.com/Jagrmi-C/gostarted/project/logger"
	"github.com/Jagrmi-C/gostarted/project/models"
)

func init()  {
	logger.LoggerInitialization()
}


func GetGroupHandler(w http.ResponseWriter, req *http.Request) {
	uuid := mux.Vars(req)["uuid"]

	var group models.Group
	err := db.GetGroup(uuid, &group)
	if err != nil {
		lr.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	lr.Info("Get group from DB:", group.UUID)

	err = json.NewEncoder(w).Encode(group)
	if err != nil {
		lr.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func GetGroupsHandler(w http.ResponseWriter, req *http.Request) {
	groups, err := db.GetGroups()

	var bodyStruct models.GroupsStruct

	bodyStruct.Groups = groups
	if err != nil {
		lr.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = json.NewEncoder(w).Encode(bodyStruct)
	if err != nil {
		lr.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func UpdateGroupHandler(w http.ResponseWriter, req *http.Request) {
	uuid := mux.Vars(req)["uuid"]

	var group models.Group
	err := json.NewDecoder(req.Body).Decode(&group)
    if err != nil {
		lr.Error(err)
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
	}

	group.UUID = uuid
	group.DT = helpers.GetCurrentLocalTime()

	err = db.UpdateGroup(&group)
	lr.Info("Update group:", group)
	if err != nil {
		lr.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(group)
	if err != nil {
		lr.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func CreateGroupHandler(w http.ResponseWriter, req *http.Request) {
	var group models.Group
	err := json.NewDecoder(req.Body).Decode(&group)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
	}

	err = db.CreateGroup(&group)
	lr.Info("Create group:", group)

	if err != nil {
		lr.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(group)

	if err != nil {
		lr.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func DeleteGroupHandler(w http.ResponseWriter, req *http.Request) {
	uuid := mux.Vars(req)["uuid"]

	err := db.DeleteGroup(uuid)
	if err != nil {
		lr.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	lr.Info("Delete group from DB with uuid:", uuid)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
