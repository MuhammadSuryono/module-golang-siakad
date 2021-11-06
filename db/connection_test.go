package db

import (
	"testing"

	"github.com/joho/godotenv"
)

func TestConnectionMysql(t *testing.T) {
	godotenv.Load()
	InitConnectionFromEnvirontment().CreateNewConnection()
}
