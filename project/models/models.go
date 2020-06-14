package models

import "time"

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

// Structure for create slice with groups
type TasksStruct struct {
    Task	[]Task
}

// Group shema of the groups table
type Group struct {
	UUID        string 	    `json:"uuid"`
	Title       string 	    `json:"title"`
	DT          time.Time 	`json:"dt"`
}

// Structure for create slice with groups
type GroupsStruct struct {
    Groups	[]Group
}

type TimeFrame struct {
	UUID        string 	    `json:"uuid"`
	TaskUUID    string 	    `json:"task_uuid"`
    FROM        time.Time 	`json:"from"`
    TO          time.Time 	`json:"to"`
}

func CreateTask(title string) *Task {
	exampleUUID := "5f3292fd-3238-444e-96d2-ee313d140166"
    return &Task{exampleUUID, title, exampleUUID}  // enforce the default value here
}
