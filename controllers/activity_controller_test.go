package controllers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetAllActivitiesWithInvalidID(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/activities", nil)
	r.SetPathValue("userID", "s")

	w := httptest.NewRecorder()

	activity_controller.AllActivities(w, r)

	if w.Code == http.StatusOK {
		t.Errorf("unexpected status code %d", w.Code)
	}

}

func TestGetAllActivities(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/activities", nil)
	r.SetPathValue("userID", "2")

	w := httptest.NewRecorder()

	activity_controller.AllActivities(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("unexpected status code %d", w.Code)
	}

	response, err := io.ReadAll(w.Result().Body)

	if err != nil {
		t.Errorf("unexpected error reading response body: %s", err.Error())
	}

	if !strings.Contains(string(response), "Hiking") || !strings.Contains(string(response), "Lifting") {
		t.Errorf("response does not contain both hiking and lifting: %s", string(response))
	}

}

func TestAddActivityInvalidIntensity(t *testing.T) {
	r := httptest.NewRequest(http.MethodPost, "/activities", strings.NewReader("{\"name\": \"Running\",\"intensity\": \"super high\",\"duration\": 20,\"date\": \"2024-04-07T00:00:00Z\"}"))
	r.SetPathValue("userID", "1")
	w := httptest.NewRecorder()
	activity_controller.AddActivity(w, r)
	if w.Code != http.StatusBadRequest {
		t.Errorf("unexpected repsonse code. Expected %d Recieved %d", http.StatusBadRequest, w.Code)
	}
}
