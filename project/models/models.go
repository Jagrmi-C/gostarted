package models

import (
	"time"
)

// User schema of the user table
type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
	Age      int64  `json:"age"`
}

// Task shema of the tasks table
type Task struct {
	UUID      string `json:"uuid"`
	Title     string `json:"title"`
	GroupUUID string `json:"group_uuid"`
}

// Group shema of the groups table
type Group struct {
	UUID  string    `json:"uuid"`
	Title string    `json:"title"`
	DT    time.Time `json:"dt"`
}

// TimeFrame shema of the timeframes table
type TimeFrame struct {
	UUID     string    `json:"uuid"`
	TaskUUID string    `json:"task_uuid"`
	FROM     time.Time `json:"from"`
	TO       time.Time `json:"to"`
}

// Structure TimeFrame for a task slice
type TaskTimeFrame struct {
	FROM time.Time `json:"from"`
	TO   time.Time `json:"to"`
}

// Structure for task slice with timeframes
type TaskInformation struct {
	UUID       string          `json:"uuid"`
	Title      string          `json:"title"`
	GroupUUID  string          `json:"group_uuid"`
	TimeFrames []TaskTimeFrame `json:"time_frames"`
}

// Structure for group slice with tasks
type GroupInformation struct {
	UUID  string `json:"uuid"`
	Title string `json:"title"`
	Tasks []Task `json:"tasks"`
}
