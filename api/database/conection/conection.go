package conection

import (
	"api_pattern_go/api/global"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDatabase() {
	DB, err = gorm.Open(postgres.Open(getStringConection()), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	log.Println("Connected to database")
}

func MakeMigrations() {
	if err = DB.AutoMigrate(global.GetModelsList()...); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
}

func getStringConection() string {
	dns := "host=" + os.Getenv("DB_HOST") +
		" user=" + os.Getenv("DB_USER") +
		" password=" + os.Getenv("DB_PASSWORD") +
		" dbname=" + os.Getenv("DB_NAME") +
		" port=" + os.Getenv("DB_PORT") +
		" sslmode=disable"
	return dns
}

func GetConnection() (*gorm.DB, error) {
	DB, err = gorm.Open(postgres.Open(getStringConection()), &gorm.Config{})
	if err != nil {
		return DB, fmt.Errorf("Could not connect to database: %v", err)
	}
	return DB, nil
}
