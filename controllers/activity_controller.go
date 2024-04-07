package controllers

import (
	"fmt"
	"net/http"
	"strconv"
)

func AddActivity(w http.ResponseWriter, r *http.Request) {
	//add an activity for a user
	userIDString := r.PathValue("userID")
	userID, _ := strconv.Atoi(userIDString)
	res := fmt.Sprintf("added an activity for user %d", userID)
	w.Write([]byte(res))
}

func AllActivities(w http.ResponseWriter, r *http.Request) {
	//get activity list for a user
	userIDString := r.PathValue("userID")
	userID, _ := strconv.Atoi(userIDString)
	res := fmt.Sprintf("all activities for user %d", userID)
	w.Write([]byte(res))
}

func ActivityStats(w http.ResponseWriter, r *http.Request) {
	//get needed stats for the user
	userIDString := r.PathValue("userID")
	userID, _ := strconv.Atoi(userIDString)
	res := fmt.Sprintf("activity stats for user %d", userID)
	w.Write([]byte(res))

}

func ActivityHeatMap(w http.ResponseWriter, r *http.Request) {
	//return data needed for heatmap
	userIDString := r.PathValue("userID")
	userID, _ := strconv.Atoi(userIDString)
	res := fmt.Sprintf("activity heatmap for user %d", userID)
	w.Write([]byte(res))
}
