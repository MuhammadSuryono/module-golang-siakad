package db

import (
	"testing"

	"github.com/joho/godotenv"
)

type Tes struct {
	ID   int64
	Name string
}

func TestConnectionMysql(t *testing.T) {
	godotenv.Load()
	InitConnection("localhost", "3306", "root", "Mamang_Sekayu.97", "db_class", "mysql").CreateNewConnection()

	Connection.AutoMigrate(&Tes{})
}
