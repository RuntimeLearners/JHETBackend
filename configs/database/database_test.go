package database_test

import (
	"JHETBackend/configs/database"
	"log"
	"testing"
)

func TestDBInit(t *testing.T) {
	db := database.DataBase
	log.Printf("%v", db)
}
