package database

import (
	"database/sql"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/idugan100/fitness_server_330/models"
)

var connection *sql.DB
var activity_repo ActivityRepository
var notification_repo NotificationRepository
var user_repo UserRepository

func TestMain(m *testing.M) {
	connection, _ = Connect("/Users/isaacdugan/code/fitness_server_330/database/test.db")
	activity_repo = NewActivityRepository(connection)
	user_repo = NewUserRepository(connection)
	notification_repo = NewNotificationRepository(connection, user_repo)

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

func TestActivitiesByUserId(t *testing.T) {
	modelList, err := activity_repo.AllByUserId(1)

	if err != nil {
		t.Errorf("unexpected error %s", err.Error())
	}

	if len(modelList) != 2 {
		t.Errorf("unexpected number of results. Expected 2 found %d", len(modelList))
	}

	if modelList[0].Intensity != "Medium" || modelList[1].Intensity != "High" {
		t.Errorf("unexpected intensity results")
	}

	if modelList[0].Duration != 20 || modelList[1].Duration != 20 {
		t.Errorf("unexpected duration results")
	}

	if modelList[0].Name != "Running" || modelList[1].Name != "Swimming" {
		t.Errorf("unexpected duration results")
	}
}

func TestGroupStats(t *testing.T) {
	stats, err := activity_repo.GroupStats()

	if err != nil {
		t.Errorf("unexpected error %s", err.Error())
	}

	if stats.TotalHigh != 40 || stats.TotalLow != 20 || stats.TotalMedium != 60 {
		t.Errorf("unexpected intensity totals. Low: %d, Medium: %d, High: %d", stats.TotalLow, stats.TotalMedium, stats.TotalHigh)
	}

	if len(stats.Days) != 1 {
		t.Errorf("unexpected days count %d", len(stats.Days))
	}

	if stats.DaysExercised != 1 {
		t.Errorf("unexpected number of days exercised count %d", stats.DaysExercised)
	}
}

func TestUserStats(t *testing.T) {
	stats, err := activity_repo.UserStats(2)

	if err != nil {
		t.Errorf("unexpected error %s", err.Error())
	}

	if stats.TotalHigh != 0 || stats.TotalLow != 20 || stats.TotalMedium != 20 {
		t.Errorf("unexpected intensity totals. Low: %d, Medium: %d, High: %d", stats.TotalLow, stats.TotalMedium, stats.TotalHigh)
	}

	if len(stats.Days) != 1 {
		t.Errorf("unexpected days count %d", len(stats.Days))
	}

	if stats.DaysExercised != 1 {
		t.Errorf("unexpected number of days exercised count %d", stats.DaysExercised)
	}
}

func TestCreateActivity(t *testing.T) {

	a := models.Activity{Name: "Boxing", Intensity: "High", Duration: 40, Date: time.Now(), UserID: 3}
	err := activity_repo.Create(a)
	if err != nil {
		t.Errorf("unexpected error %s", err.Error())
	}

	activityList, err := activity_repo.AllByUserId(3)
	if err != nil {
		t.Errorf("unexpected error %s", err.Error())
	}

	if len(activityList) != 3 {
		t.Errorf("unexpected number of results expected 3, got %d", len(activityList))
	}

}
