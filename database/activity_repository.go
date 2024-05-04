package database

import (
	"database/sql"

	"github.com/idugan100/fitness_server_330/models"
)

type ActivityRepository struct {
	Connection *sql.DB
}

func NewActivityRepository(conn *sql.DB) ActivityRepository {
	return ActivityRepository{Connection: conn}
}

func (a ActivityRepository) AllByUserId(userId int) ([]models.Activity, error) {

	var activityList []models.Activity
	var activity models.Activity

	rows, err := a.Connection.Query("SELECT * FROM Activities WHERE userID=? ORDER BY id DESC", userId)
	if err != nil {
		return activityList, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&activity.Id, &activity.Name, &activity.UserID, &activity.Duration, &activity.Intensity, &activity.Date)
		if err != nil {
			return activityList, err
		}
		activityList = append(activityList, activity)
	}

	return activityList, nil
}

func (a ActivityRepository) UserStats(userId int) (models.UserStats, error) {
	var stats models.UserStats

	rows, err := a.Connection.Query("SELECT COALESCE(COUNT(DISTINCT date(date)),0) FROM Activities WHERE userID=?", userId)
	if err != nil {
		return stats, err
	}
	rows.Next()
	err = rows.Scan(&stats.DaysExercised)
	if err != nil {
		return stats, err
	}
	rows.Close()

	//total high
	rows, err = a.Connection.Query("SELECT COALESCE(SUM(duration),0) FROM Activities WHERE userID=? AND intensity=?", userId, "High")
	if err != nil {
		return stats, err
	}
	if rows.Next() {
		err = rows.Scan(&stats.TotalHigh)
		if err != nil {
			return stats, err
		}
	}
	rows.Close()

	//total medium
	rows, err = a.Connection.Query("SELECT COALESCE(SUM(duration),0) FROM Activities WHERE userID=? AND intensity=?", userId, "Medium")
	if err != nil {
		return stats, err
	}
	if rows.Next() {
		err = rows.Scan(&stats.TotalMedium)
		if err != nil {
			return stats, err
		}
	}
	rows.Close()

	//total low
	rows, err = a.Connection.Query("SELECT COALESCE(SUM(duration),0) FROM Activities WHERE userID=? AND intensity=?", userId, "Low")
	if err != nil {
		return stats, err
	}
	if rows.Next() {
		err = rows.Scan(&stats.TotalLow)
		if err != nil {
			return stats, err
		}
	}
	rows.Close()

	//heatmap
	rows, err = a.Connection.Query("SELECT COALESCE(SUM(duration),0), date(date) FROM Activities WHERE userID=? GROUP BY date(date)", userId)
	if err != nil {
		return stats, err
	}
	var daily models.DayActivityTotal
	for rows.Next() {
		err = rows.Scan(&daily.TotalMinutes, &daily.Date)
		if err != nil {
			return stats, err
		}
		stats.Days = append(stats.Days, daily)
	}
	rows.Close()
	return stats, nil
}

func (a ActivityRepository) GroupStats() (models.UserStats, error) {
	var stats models.UserStats

	rows, err := a.Connection.Query("SELECT COALESCE(COUNT(DISTINCT date(date)),0) FROM Activities")
	if err != nil {
		return stats, err
	}
	rows.Next()
	err = rows.Scan(&stats.DaysExercised)
	if err != nil {
		return stats, err
	}
	rows.Close()

	//total high
	rows, err = a.Connection.Query("SELECT COALESCE(SUM(duration),0) FROM Activities WHERE intensity=?", "High")
	if err != nil {
		return stats, err
	}
	if rows.Next() {
		err = rows.Scan(&stats.TotalHigh)
		if err != nil {
			return stats, err
		}
	}
	rows.Close()

	//total medium
	rows, err = a.Connection.Query("SELECT COALESCE(SUM(duration),0) FROM Activities WHERE intensity=?", "Medium")
	if err != nil {
		return stats, err
	}
	if rows.Next() {
		err = rows.Scan(&stats.TotalMedium)
		if err != nil {
			return stats, err
		}
	}
	rows.Close()

	//total low
	rows, err = a.Connection.Query("SELECT COALESCE(SUM(duration),0) FROM Activities WHERE intensity=?", "Low")
	if err != nil {
		return stats, err
	}
	if rows.Next() {
		err = rows.Scan(&stats.TotalLow)
		if err != nil {
			return stats, err
		}
	}
	rows.Close()

	//heatmap
	rows, err = a.Connection.Query("SELECT COALESCE(SUM(duration),0), date(date) FROM Activities GROUP BY date(date)")
	if err != nil {
		return stats, err
	}
	var daily models.DayActivityTotal
	for rows.Next() {
		err = rows.Scan(&daily.TotalMinutes, &daily.Date)
		if err != nil {
			return stats, err
		}
		stats.Days = append(stats.Days, daily)
	}
	rows.Close()
	return stats, nil
}

func (a ActivityRepository) Create(act models.Activity) error {
	_, err := a.Connection.Exec("INSERT INTO Activities (name, userId,duration,intensity, date) VALUES (?,?,?,?,?)", act.Name, act.UserID, act.Duration, act.Intensity, act.Date.Format("2006-01-02"))
	if err != nil {
		return err
	}
	return nil
}
