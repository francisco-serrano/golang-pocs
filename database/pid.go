package database

import (
	"fmt"
)

func RunPid() {
	runPid()
}

func runPid() {
	db, err := initializeDatabaseConnection()
	if err != nil {
		panic(err)
	}

	statusArr := []JobStatus{StatusRunning, StatusTerminatedFailure}

	fmt.Println(statusArr)

	db.LogMode(true)

	var job Job
	if err := db.Where(`id = ?`, uint32(618623)).Find(&job).Error; err != nil {
		panic(err)
	}

	fmt.Println(job)
}
