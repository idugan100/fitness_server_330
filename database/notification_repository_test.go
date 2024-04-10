package database

import (
	"testing"

	"github.com/idugan100/fitness_server_330/models"
)

func TestNotificationsByUserId(t *testing.T) {
	res, err := notification_repo.ByUserId(2)
	if err != nil {
		t.Errorf("unexpected error when getting activities for a user: %s", err.Error())
	}
	if len(res) != 2 {
		t.Errorf("unexpected number of notifications found. Expected 2 got %d", len(res))
	}

	expected_message := []string{
		"welcome to the app", "remember to workout today!",
	}
	for i, notificaton := range res {
		if notificaton.Message != expected_message[i] {
			t.Errorf("unexpected notification found: %s", notificaton.Message)
		}
	}
}

func TestReadNotification(t *testing.T) {
	err := notification_repo.Read(1, 1)
	if err != nil {
		t.Errorf("unexpected error when reading unread notification: %s", err.Error())
	}
	res, _ := notification_repo.ByUserId(1)

	for _, notification := range res {
		if notification.Id == 1 && notification.Read == false {
			t.Errorf("notification %d not read ", notification.Id)
		}
	}

	err = notification_repo.Read(6, 3)
	if err != nil {
		t.Errorf("unexpected error when reading already reads notification: %s", err.Error())
	}
	for _, notification := range res {
		if notification.Id == 6 && notification.Read == false {
			t.Errorf("notification %d not read ", notification.Id)
		}
	}
}

func TestCreateNotification(t *testing.T) {
	notification := models.Notification{Message: "new message"}
	err := notification_repo.Create(notification)
	if err != nil {
		t.Errorf("unexpected error when creating notification: %s", err.Error())
	}

	userid_list := []int{1, 2, 3}

	for _, i := range userid_list {
		res, _ := notification_repo.ByUserId(i)
		if len(res) != 3 {
			t.Errorf("notification not properly created for user id %d", i)
		}
	}
}

func TestDeleteNotification(t *testing.T) {
	err := notification_repo.Delete(2, 2)
	if err != nil {
		t.Errorf("unexpected error deleting notification: %s", err.Error())
	}
	res, _ := notification_repo.ByUserId(2)
	if len(res) != 2 {
		t.Errorf("deletion did not work as expected. Expected 1 notification found %d", len(res))
	}
}
