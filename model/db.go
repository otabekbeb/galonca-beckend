package model

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"net/url"
	"os"
)

var db *gorm.DB

func init() {
	e := godotenv.Load()
	if e != nil {
		log.Fatalf("Error read enveroment %v", e)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	dbUri := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=%s", username, password, dbHost, dbName, url.QueryEscape("Asia/Almaty"))
	log.Println(dbUri)
	conn, err := gorm.Open(mysql.Open(dbUri), &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Info),
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		fmt.Print(err)
	}
	db = conn
	err = db.AutoMigrate(
		TslUser{},
		Cargo{},
		Truck{},
		TransportType{},
		LoadingType{},
		Addition{},
		ServiceStation{},
		ServiceCategory{},
		SparePart{},
		Brand{},
		Transport{},
		RoadsideService{},
		Company{},
		Favorite{},
		Like{},
		CargoType{},
	)
	if err != nil {
		log.Fatalf("Can not migrate %v\n", err)
	}
}

func GetDB() *gorm.DB {
	return db
}
