package database

import "time"

type JobStatus uint8

const (
	StatusPendingToRun JobStatus = iota
	StatusRunning
	StatusTerminatedSuccess
	StatusTerminatedFailure
	StatusAwaitingDeviceIngester
	StatusAwaitingLookAlike
	StatusAwaitingDevicer
	StatusAwaitingGeneric
	StatusBackup JobStatus = 10
)

type Job struct {
	Id             uint32     `gorm:"column:id"`
	ClientId       uint16     `gorm:"column:id_client"`
	SegmentId      *int       `gorm:"column:id_segment"`
	Params         string     `gorm:"column:params"`
	Response       *string    `gorm:"column:response"`
	Internal       *string    `gorm:"column:internal"`
	Type           string     `gorm:"column:type"`
	Created        time.Time  `gorm:"column:created"`
	Finished       *time.Time `gorm:"column:finished"`
	Progress       int8       `gorm:"column:progress"`
	Pid            uint32     `gorm:"column:pid"`
	Status         JobStatus  `gorm:"column:status"`
	Started        time.Time  `gorm:"column:started"`
	Priority       uint16     `gorm:"column:priority"`
	Description    string     `gorm:"column:description"`
	Queue          string     `gorm:"column:queue"`
	DependsOnJobId *int       `gorm:"column:depends_on_job_id"`
	BackupParams   string     `gorm:"column:backup_params"`
	Monthly        int8       `gorm:"column:monthly"`
	Who            string     `gorm:"column:who"`
	Retries        int16      `gorm:"column:retries"`
}

func (Job) TableName() string {
	return "job"
}
