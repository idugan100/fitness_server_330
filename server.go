package main

import (
	"fmt"
	"net/http"

	"github.com/idugan100/fitness_server_330/controllers"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("response"))
}
func main() {
	mux := http.NewServeMux()
	// notification routes
	mux.HandleFunc("GET /notifications/{userID}", controllers.SetUserMiddleware(controllers.AllNotifications))
	mux.HandleFunc("GET /notification/read/{userID}/{notificationID}", controllers.SetUserMiddleware(controllers.ReadNotification))
	mux.HandleFunc("GET /notification/delete/{userID}/{notificationID}", controllers.SetUserMiddleware(controllers.DeleteNotification))
	mux.HandleFunc("POST /notification/create", controllers.SetUserMiddleware(controllers.CreateNotification))
	//activity routes
	mux.HandleFunc("GET /activities/{userID}", controllers.SetUserMiddleware(controllers.AllActivities))
	mux.HandleFunc("GET /activities/stats/{userID}", controllers.SetUserMiddleware(controllers.ActivityStats))
	mux.HandleFunc("GET /activities/heatmap/{userID}", controllers.SetUserMiddleware(controllers.ActivityHeatMap))
	mux.HandleFunc("POST /activties/{userID}", controllers.SetUserMiddleware(controllers.AddActivity))
	//auth routes
	mux.HandleFunc("POST /signup", controllers.Signup)
	mux.HandleFunc("POST /login", controllers.Login)
	mux.HandleFunc("GET /logout", controllers.SetUserMiddleware(controllers.Logout))

	s := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	fmt.Println("server starting ...")
	panic(s.ListenAndServe())
}
