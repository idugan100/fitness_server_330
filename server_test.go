package main

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/idugan100/fitness_server_330/controllers"
	"github.com/idugan100/fitness_server_330/database"
	"github.com/idugan100/fitness_server_330/models"
)

var connection *sql.DB

type RequestTest struct {
	Method string
	Path   string
	Body   io.Reader
	Code   int
}

var RequestTable = []RequestTest{
	//activities
	{Method: "GET", Path: "/activities/stats/all", Body: nil, Code: http.StatusOK},
	{Method: "GET", Path: "/activities/stats/1", Body: nil, Code: http.StatusOK},
	{Method: "GET", Path: "/activities/1", Body: nil, Code: http.StatusOK},
	{Method: "POST", Path: "/activities/1", Body: strings.NewReader("{\"name\": \"Running\",\"intensity\": \"Medium\",\"duration\": 20,\"date\": \"2024-04-07T00:00:00Z\"}"), Code: http.StatusCreated},

	//users
	{Method: "GET", Path: "/allusers", Body: nil, Code: http.StatusOK},
	//successful login
	{Method: "POST", Path: "/login", Body: strings.NewReader("{\"username\":\"tom\",\"password\":\"123\"}"), Code: http.StatusOK},
	//incorrect password
	{Method: "POST", Path: "/login", Body: strings.NewReader("{\"username\":\"tom\",\"password\":\"12\"}"), Code: http.StatusBadRequest},
	//duplicate username
	{Method: "POST", Path: "/signup", Body: strings.NewReader("{\"username\":\"tom\",\"password\":\"123\"}"), Code: http.StatusBadRequest},
	//successful signup
	{Method: "POST", Path: "/signup", Body: strings.NewReader("{\"username\":\"tommy\",\"password\":\"123\"}"), Code: http.StatusCreated},

	//notifications
	{Method: "GET", Path: "/notifications/1", Body: nil, Code: http.StatusOK},
	{Method: "GET", Path: "/notifications/read/1/1", Body: nil, Code: http.StatusOK},
	{Method: "GET", Path: "/notifications/delete/2/2", Body: nil, Code: http.StatusOK},
	{Method: "POST", Path: "/notifications", Body: strings.NewReader("{\"message\":\"this is a test message\"}"), Code: http.StatusCreated},

	//documentation
	{Method: "GET", Path: "/", Body: nil, Code: http.StatusOK},
}

func TestMain(m *testing.M) {
	connection, _ = database.Connect("/Users/isaacdugan/code/fitness_server_330/database/test.db")

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

func TestRoutes(t *testing.T) {
	//get an admin token
	server := SetupServer(connection, "8080")
	token, err := controllers.CreateToken(models.User{Id: 1, IsAdmin: true})
	if err != nil {
		t.Errorf("unexpected error when creating admin token %s", err.Error())
	}

	for _, test := range RequestTable {
		r := httptest.NewRequest(test.Method, test.Path, test.Body)
		r.Header.Add("Authorization", "Bearer "+token)
		w := httptest.NewRecorder()

		server.Handler.ServeHTTP(w, r)

		if w.Code != test.Code {
			t.Errorf("unexpected status code when performing a %s request to %s. Expected %d, Received %d", test.Method, test.Path, test.Code, w.Code)
		}
	}

}
