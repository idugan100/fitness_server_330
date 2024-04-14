package controllers

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/idugan100/fitness_server_330/database"
)

var n_controller NotificationController
var auth_controller AuthController
var user_repo database.UserRepository
var notification_repo database.NotificationRepository
var activity_controller ActivityController

func TestMain(m *testing.M) {
	connection, _ := database.Connect("/Users/isaacdugan/code/fitness_server_330/database/test.db")
	user_repo = database.NewUserRepository(connection)
	auth_controller = NewAuthController(user_repo)
	notification_repo = database.NewNotificationRepository(connection, user_repo)
	n_controller = CreateNotificationController(notification_repo)
	activity_repo := database.NewActivityRepository(connection)
	activity_controller = NewActivityController(activity_repo)

	clear, _ := os.ReadFile("/Users/isaacdugan/code/fitness_server_330/database/clear.sql")
	schema, _ := os.ReadFile("/Users/isaacdugan/code/fitness_server_330/database/schema.sql")
	seed, _ := os.ReadFile("/Users/isaacdugan/code/fitness_server_330/database/seed.sql")

	connection.Exec(string(clear))
	connection.Exec(string(schema))
	connection.Exec(string(seed))

	fmt.Println("setup complete")

	code := m.Run()

	connection.Exec(string(clear))
	connection.Close()
	fmt.Println("tear down complete")

	os.Exit(code)
}

func TestAllNotifications(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/notifications/", nil)
	r.SetPathValue("userID", "2")
	w := httptest.NewRecorder()
	n_controller.AllNotifications(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("unexpected status code. Expected %d. Recieved %d.", http.StatusOK, w.Code)
	}

	responseString, err := io.ReadAll(w.Result().Body)
	fmt.Print(responseString)
	if err != nil {
		t.Errorf("unexpected error %s", err.Error())
	}

	if !strings.Contains(string(responseString), "welcome to the app") || !strings.Contains(string(responseString), "remember to workout today!") {
		t.Errorf("json response does not contain expected output. Output: %s", responseString)
	}
}

func TestCreateValidNotification(t *testing.T) {
	r := httptest.NewRequest(http.MethodPost, "/notifications", strings.NewReader("{\"message\":\"test notification\"}"))
	w := httptest.NewRecorder()

	n_controller.CreateNotification(w, r)
	if w.Code != http.StatusCreated {
		t.Errorf("unexpected response code. Expected: %d Recieved: %d", http.StatusCreated, w.Code)
	}

}

func TestCreateInValidNotification(t *testing.T) {
	r := httptest.NewRequest(http.MethodPost, "/notifications", strings.NewReader("{\"text\":\"test notification\"}"))
	w := httptest.NewRecorder()

	n_controller.CreateNotification(w, r)
	if w.Code != http.StatusBadRequest {
		t.Errorf("unexpected response code. Expected: %d Recieved: %d", http.StatusBadRequest, w.Code)
	}
}

func TestCreateInValidJSONNotification(t *testing.T) {
	r := httptest.NewRequest(http.MethodPost, "/notifications", strings.NewReader("{text\":\"test notification\"}"))
	w := httptest.NewRecorder()

	n_controller.CreateNotification(w, r)
	if w.Code != http.StatusBadRequest {
		t.Errorf("unexpected response code. Expected: %d Recieved: %d", http.StatusBadRequest, w.Code)
	}
}

func TestUnReadReadNotification(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/notifications/read", nil)
	r.SetPathValue("userID", "1")
	r.SetPathValue("notificationID", "1")
	w := httptest.NewRecorder()
	n_controller.ReadNotification(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("unexpected status code. Expected: %d Recieved: %d", http.StatusOK, w.Code)
	}

	notifications, err := notification_repo.ByUserId(1)

	if err != nil {
		t.Errorf("unexpected error  quering notifications: %s", err.Error())
	}

	for _, n := range notifications {
		if n.Id == 1 && n.Read == false {
			t.Errorf("notification not marked read in database")
		}
	}
}

func TestDeleteNotification(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/notifications/delete", nil)
	r.SetPathValue("userID", "1")
	r.SetPathValue("notificationID", "1")
	w := httptest.NewRecorder()
	n_controller.DeleteNotification(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("unexpected status code. Expected: %d Recieved: %d", http.StatusOK, w.Code)
	}

	notifications, err := notification_repo.ByUserId(1)

	if err != nil {
		t.Errorf("unexpected error  quering notifications: %s", err.Error())
	}

	if len(notifications) != 2 {
		t.Errorf("unexpected number of notifications for user. Expected %d Revieved %d", 2, len(notifications))
	}
}
