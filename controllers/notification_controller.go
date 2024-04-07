package controllers

import (
	"fmt"
	"net/http"
	"strconv"
)

func AllNotifications(w http.ResponseWriter, r *http.Request) {

	userIDString := r.PathValue("userID")
	userID, _ := strconv.Atoi(userIDString)
	res := fmt.Sprintf("all notifications for user %d", userID)
	w.Write([]byte(res))
}

func ReadNotification(w http.ResponseWriter, r *http.Request) {
	userIDString := r.PathValue("userID")
	userID, _ := strconv.Atoi(userIDString)

	notificationIDString := r.PathValue("notificationID")
	notificationID, _ := strconv.Atoi(notificationIDString)

	//mark notification as read here

	res := fmt.Sprintf("read notification %d for user %d", notificationID, userID)
	w.Write([]byte(res))
}

func DeleteNotification(w http.ResponseWriter, r *http.Request) {
	userIDString := r.PathValue("userID")
	userID, _ := strconv.Atoi(userIDString)

	notificationIDString := r.PathValue("notificationID")
	notificationID, _ := strconv.Atoi(notificationIDString)

	//delete notification here
	res := fmt.Sprintf("deleted notification %d for user %d", notificationID, userID)
	w.Write([]byte(res))
}

func CreateNotification(w http.ResponseWriter, r *http.Request) {
	//get message from form or path
	//insert notification for each user
	w.Write([]byte("notification created and sent to each user"))
}
