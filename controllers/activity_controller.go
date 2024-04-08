package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/idugan100/fitness_server_330/database"
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
	res := fmt.Sprintf("added an activity for user %d", userID)
	w.Write([]byte(res))
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
	}
	data, _ := json.Marshal(&stats)
	w.Write(data)

}

func (a ActivityController) GroupActivityStats(w http.ResponseWriter, r *http.Request) {

	stats, err := a.ActivityRepo.GroupStats()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	data, _ := json.Marshal(&stats)
	w.Write(data)

}
