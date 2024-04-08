package database

import (
	"testing"
)

func TestConnection(t *testing.T) {
	_, err := Connect("/Users/isaacdugan/code/fitness_server_330/database/database.db")
	if err != nil {
		t.Errorf("error connecting to databse %s", err.Error())
	}
	_, err = Connect("/Users/isaacdugan/code/fitness_server_330/database/test.db")
	if err != nil {
		t.Errorf("error connecting to databse %s", err.Error())
	}
}
