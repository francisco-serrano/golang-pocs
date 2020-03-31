package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func initializeDatabaseConnection() (*gorm.DB, error) {
	user := "root"
	pass := "root"
	host := "localhost"
	port := "3306"
	name := "rely"

	connectionURL := fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=true", user, pass, host, port, name)

	db, err := gorm.Open("mysql", connectionURL)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func Run() {
	run()
}

func run() {
	db, err := initializeDatabaseConnection()
	if err != nil {
		panic(err)
	}

	statusArr := []JobStatus{StatusRunning, StatusTerminatedFailure}

	fmt.Println(statusArr)

	db.LogMode(true)

	var jobs []Job
	if err := db.Where(`status in (?) AND who != ''`, statusArr).Find(&jobs).Error; err != nil {
		panic(err)
	}

	fmt.Println(len(jobs))
}
