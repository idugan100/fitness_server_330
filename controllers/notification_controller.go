package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/idugan100/fitness_server_330/database"
)

type NotificationController struct {
	NotificationRepo database.NotificationRepository
}

func CreateNotificationController(NRepo database.NotificationRepository) NotificationController {
	return NotificationController{NotificationRepo: NRepo}
}

func (n NotificationController) AllNotifications(w http.ResponseWriter, r *http.Request) {

	userIDString := r.PathValue("userID")
	userID, _ := strconv.Atoi(userIDString)
	notificationList, err := n.NotificationRepo.ByUserId(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	json, err := json.Marshal(notificationList)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write(json)

}

func (n NotificationController) ReadNotification(w http.ResponseWriter, r *http.Request) {
	userIDString := r.PathValue("userID")
	userID, _ := strconv.Atoi(userIDString)

	notificationIDString := r.PathValue("notificationID")
	notificationID, _ := strconv.Atoi(notificationIDString)

	err := n.NotificationRepo.Read(notificationID, userID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (n NotificationController) DeleteNotification(w http.ResponseWriter, r *http.Request) {
	userIDString := r.PathValue("userID")
	userID, _ := strconv.Atoi(userIDString)

	notificationIDString := r.PathValue("notificationID")
	notificationID, _ := strconv.Atoi(notificationIDString)

	err := n.NotificationRepo.Delete(notificationID, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (n NotificationController) CreateNotification(w http.ResponseWriter, r *http.Request) {
	//get message from form or path
	//insert notification for each user
	w.Write([]byte("notification created and sent to each user"))
}
