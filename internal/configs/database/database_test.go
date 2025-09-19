package database_test

import (
	"JHETBackend/internal/configs/database"
	"testing"
)

func TestDBInit(t *testing.T) {
	database.InitDatabase()
}
