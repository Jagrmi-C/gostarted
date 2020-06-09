package db

import (
	"context"
	"fmt"
	"os"
	// "time"

	"github.com/Jagrmi-C/gostarted/project/models"
	"github.com/jackc/pgx/v4"
)


const (
	GetTaskQuery    = "SELECT uuid, title FROM tasks WHERE uuid=$1"
	// GetTasksQuery   = "SELECT uuid, title, dt FROM tasks LIMIT $1 OFFSET $2"
	GetTasksQuery   = "SELECT uuid, title FROM tasks"
	UpdateTaskQuery = "UPDATE tasks SET title=$1,group_uuid=$2 WHERE uuid=$3"
	UpdateTaskQueryIfTitle = "UPDATE tasks SET title=$1 WHERE uuid=$2"
	UpdateTaskQueryIfGroup = "UPDATE tasks SET group_uuid=$1 WHERE uuid=$2"
	DeleteTaskQuery = "DELETE FROM tasks WHERE uuid=$1"
	CreateTaskQuery = "INSERT INTO tasks(title, group_uuid) VALUES($1, $2) RETURNING uuid"
)

func CreateConnection() *pgx.Conn {
	conn, err := pgx.Connect(
		context.Background(),
		os.Getenv("DATABASE_URL"),
	)

	if err != nil {
		panic(err)
	}

	// check the connection
	err = conn.Ping(context.Background())

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	// return the connection
	return conn
}

func CheckDb() {
	fmt.Println("TEST", os.Getenv("DATABASE_URL"))
	conn, err := pgx.Connect(
		context.Background(),
		os.Getenv("DATABASE_URL"),
	)
	if err != nil {
        panic(err)
	}
	err = conn.Ping(context.Background())
    if err != nil {
        panic(err)
	}

	fmt.Println("Successfully connected!")
}

func GetTask(db *pgx.Conn, uuid string, t *models.Task) error {
	err := db.QueryRow(
		context.Background(),
		GetTaskQuery,
		uuid,
	).Scan(&t.UUID, &t.Title)

	return err
}

func GetTasks(db *pgx.Conn) ([]models.Task, error) {
	var tasks []models.Task

	rows, _ := db.Query(context.Background(), GetTasksQuery)

	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.UUID, &task.Title,)
		if err != nil {
			return tasks, err
		}
		tasks = append(tasks, task)
	}

	return tasks, rows.Err()
}

func UpdateProduct(db *pgx.Conn, t *models.Task) error {
	if (t.Title != "" && t.GroupUUID != "") {
		_, err := db.Exec(
			context.Background(),
			UpdateTaskQuery,
			t.Title,
			t.GroupUUID,
			t.UUID,
		)
		return err
	} else if (t.Title != "") {
		_, err := db.Exec(
			context.Background(),
			UpdateTaskQueryIfTitle,
			t.Title,
			t.UUID,
		)
		return err
	} else {
		_, err := db.Exec(
			context.Background(),
			UpdateTaskQueryIfGroup,
			t.GroupUUID,
			t.UUID,
		)
		return err
	}
}

func CreateTask(db *pgx.Conn, t *models.Task) error {
	err := db.QueryRow(
		context.Background(),
		CreateTaskQuery,
		t.Title,
		t.GroupUUID,
	).Scan(&t.UUID)
	return err
}
