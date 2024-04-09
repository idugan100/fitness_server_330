package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	"github.com/idugan100/fitness_server_330/controllers"
	"github.com/idugan100/fitness_server_330/database"
)

func SetupServer(conn *sql.DB, port string) *http.Server {
	mux := http.NewServeMux()

	//auth routes
	user_repo := database.NewUserRepository(conn)
	auth_controller := controllers.NewAuthController(user_repo)
	mux.HandleFunc("POST /signup", auth_controller.Signup)
	mux.HandleFunc("POST /login", auth_controller.Login)
	mux.HandleFunc("GET /allusers", controllers.SetUserMiddleware(controllers.RequireAdmin(auth_controller.AllUsers)))

	//activity routes
	activity_repo := database.NewActivityRepository(conn)
	activity_controller := controllers.NewActivityController(activity_repo)
	mux.HandleFunc("GET /activities/{userID}", controllers.SetUserMiddleware(controllers.CheckPermission(activity_controller.AllActivities)))
	mux.HandleFunc("GET /activities/stats/{userID}", controllers.SetUserMiddleware(controllers.CheckPermission(activity_controller.ActivityStats)))
	mux.HandleFunc("POST /activities/{userID}", controllers.SetUserMiddleware(controllers.CheckPermission(activity_controller.AddActivity)))
	mux.HandleFunc("GET /activities/stats/all", controllers.SetUserMiddleware(controllers.RequireAdmin(activity_controller.GroupActivityStats)))

	// notification routes
	notification_repo := database.NewNotificationRepository(conn, user_repo)
	notification_controller := controllers.CreateNotificationController(notification_repo)
	mux.HandleFunc("GET /notifications/{userID}", controllers.SetUserMiddleware(controllers.CheckPermission(notification_controller.AllNotifications)))
	mux.HandleFunc("GET /notifications/read/{userID}/{notificationID}", controllers.SetUserMiddleware(controllers.CheckPermission(notification_controller.ReadNotification)))
	mux.HandleFunc("GET /notifications/delete/{userID}/{notificationID}", controllers.SetUserMiddleware(controllers.CheckPermission(notification_controller.DeleteNotification)))
	mux.HandleFunc("POST /notifications", controllers.SetUserMiddleware(controllers.RequireAdmin(notification_controller.CreateNotification)))

	//documentation
	documentaion := template.Must(template.ParseFiles("./docs.html"))

	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) { documentaion.Execute(w, nil) })
	s := http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}
	return &s
}

func main() {
	conn, err := database.Connect("/Users/isaacdugan/code/fitness_server_330/database/database.db")
	if err != nil {
		fmt.Printf("error connecting to database %s", err.Error())
		return
	}
	defer conn.Close()

	fmt.Println("server starting ...")
	panic(SetupServer(conn, "8080").ListenAndServe())
}
