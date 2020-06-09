package models

// import (
// 	"time"
// )

// User schema of the user table
type User struct {
    ID       int64  `json:"id"`
    Name     string `json:"name"`
    Location string `json:"location"`
    Age      int64  `json:"age"`
}

// Tssk shema of the tasks table
type Task struct {
	UUID      string 	`json:"uuid"`
	Title     string 	`json:"title"`
	GroupUUID string 	`json:"group_uuid"`
}

func CreateTask(title string) *Task {
	exampleUUID := "5f3292fd-3238-444e-96d2-ee313d140166"
    return &Task{exampleUUID, title, exampleUUID}  // enforce the default value here
}
