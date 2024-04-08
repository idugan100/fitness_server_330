package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/idugan100/fitness_server_330/database"
	"github.com/idugan100/fitness_server_330/models"
)

type ActivityController struct {
	ActivityRepo database.ActivityRepository
}

func NewActivityController(ActRepo database.ActivityRepository) ActivityController {
	return ActivityController{ActivityRepo: ActRepo}
}

func (a ActivityController) AddActivity(w http.ResponseWriter, r *http.Request) {
	//add an activity for a user
	userIDString := r.PathValue("userID")
	userID, _ := strconv.Atoi(userIDString)
	var activity models.Activity
	bodyString, _ := io.ReadAll(r.Body)
	json.Unmarshal(bodyString, &activity)
	activity.UserID = userID
	if !(activity.Intensity == "High" || activity.Intensity == "Medium" || activity.Intensity == "Low") {
		http.Error(w, "intensity must be either High, Medium, or Low", http.StatusBadRequest)
		return
	}
	err := a.ActivityRepo.Create(activity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (a ActivityController) AllActivities(w http.ResponseWriter, r *http.Request) {
	//get activity list for a user
	userIDString := r.PathValue("userID")
	userID, _ := strconv.Atoi(userIDString)
	activityList, err := a.ActivityRepo.AllByUserId(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, _ := json.Marshal(activityList)
	w.Write(data)
}

func (a ActivityController) ActivityStats(w http.ResponseWriter, r *http.Request) {
	//get needed stats for the user
	userIDString := r.PathValue("userID")
	userID, _ := strconv.Atoi(userIDString)
	stats, err := a.ActivityRepo.UserStats(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data, _ := json.Marshal(&stats)
	w.Write(data)

}

func (a ActivityController) GroupActivityStats(w http.ResponseWriter, r *http.Request) {

	stats, err := a.ActivityRepo.GroupStats()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data, _ := json.Marshal(&stats)
	w.Write(data)

}
