package database

import (
	"database/sql"
	"fmt"

	"github.com/idugan100/fitness_server_330/models"
)

type NotificationRepository struct {
	Connection *sql.DB
}

func NewNotificationRepository(conn *sql.DB) NotificationRepository {
	return NotificationRepository{Connection: conn}
}

func (n NotificationRepository) ByUserId(userId int) ([]models.Notification, error) {
	rows, err := n.Connection.Query("SELECT * FROM Notifications WHERE userID=?", userId)
	defer rows.Close()
	var notificationList []models.Notification
	var notification models.Notification
	if err != nil {
		return notificationList, err
	}

	for rows.Next() {
		err = rows.Scan(&notification.Id, &notification.RecipientId, &notification.Message, &notification.Read, &notification.CreatedAt)
		if err != nil {
			return notificationList, err
		}
		notificationList = append(notificationList, notification)
	}
	return notificationList, nil

}

func (n NotificationRepository) Read(notificationID, userID int) error {
	_, err := n.Connection.Exec("UPDATE Notifications SET isRead=true WHERE id=? AND userID=?", notificationID, userID)
	if err != nil {
		fmt.Print(err.Error())
		return err
	}
	return nil
}

func (n NotificationRepository) Delete(notificationID, userID int) error {
	_, err := n.Connection.Exec("DELETE FROM Notifications WHERE id=? and userID=?", notificationID, userID)
	if err != nil {
		return err
	}
	return nil
}

func (n NotificationRepository) Create(message string) {

}
