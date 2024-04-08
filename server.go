package main

import (
	"fmt"
	"net/http"

	"github.com/idugan100/fitness_server_330/controllers"
	"github.com/idugan100/fitness_server_330/database"
)

func main() {
	conn, err := database.Connect("/Users/isaacdugan/code/fitness_server_330/database/database.db")
	if err != nil {
		_ = fmt.Errorf("error connecting to database %w", err)
		return
	}
	mux := http.NewServeMux()

	//activity routes
	activity_repo := database.NewActivityRepository(conn)
	activity_controller := controllers.NewActivityController(activity_repo)
	mux.HandleFunc("GET /activities/{userID}", controllers.SetUserMiddleware(activity_controller.AllActivities))
	mux.HandleFunc("GET /activities/stats/{userID}", controllers.SetUserMiddleware(activity_controller.ActivityStats))
	mux.HandleFunc("POST /activities/{userID}", controllers.SetUserMiddleware(activity_controller.AddActivity)) //not done
	mux.HandleFunc("GET /activities/stats/all", controllers.SetUserMiddleware(activity_controller.GroupActivityStats))

	//auth routes
	user_repo := database.NewUserRepository(conn)
	auth_controller := controllers.NewAuthController(user_repo)
	mux.HandleFunc("POST /signup", auth_controller.Signup)
	mux.HandleFunc("POST /login", auth_controller.Login)
	mux.HandleFunc("GET /logout", controllers.SetUserMiddleware(auth_controller.Logout))
	mux.HandleFunc("GET /allusers", controllers.SetUserMiddleware(auth_controller.AllUsers))

	// notification routes
	notification_repo := database.NewNotificationRepository(conn, user_repo)
	notification_controller := controllers.CreateNotificationController(notification_repo)
	mux.HandleFunc("GET /notifications/{userID}", controllers.SetUserMiddleware(notification_controller.AllNotifications))
	mux.HandleFunc("GET /notifications/read/{userID}/{notificationID}", controllers.SetUserMiddleware(notification_controller.ReadNotification))
	mux.HandleFunc("GET /notifications/delete/{userID}/{notificationID}", controllers.SetUserMiddleware(notification_controller.DeleteNotification))
	mux.HandleFunc("POST /notifications/create", controllers.SetUserMiddleware(notification_controller.CreateNotification))

	s := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	fmt.Println("server starting ...")
	panic(s.ListenAndServe())
}
