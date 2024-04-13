package models

type User struct {
	Id       int    `json:"id"`
	UserName string `json:"username"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"isAdmin"`
}

type UserStats struct {
	DaysExercised int                `json:"totalDays"`
	TotalHigh     int                `json:"totalHigh"`
	TotalMedium   int                `json:"totalMedium"`
	TotalLow      int                `json:"totalLow"`
	Days          []DayActivityTotal `json:"dailyTotals"`
}

type DayActivityTotal struct {
	Date         string `json:"date"`
	TotalMinutes int    `json:"totalMinutes"`
}
