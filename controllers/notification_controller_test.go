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

var controller NotificationController

func TestMain(m *testing.M) {
	connection, _ := database.Connect("/Users/isaacdugan/code/fitness_server_330/database/test.db")
	user_repo := database.NewUserRepository(connection)
	notification_repo := database.NewNotificationRepository(connection, user_repo)
	controller = CreateNotificationController(notification_repo)

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
	r := httptest.NewRequest("GET", "/notifications/", nil)
	r.SetPathValue("userID", "2")
	w := httptest.NewRecorder()
	controller.AllNotifications(w, r)

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