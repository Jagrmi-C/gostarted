package db

import (
	"context"
	"os"

	"github.com/jackc/pgx/v4"
	lr "github.com/sirupsen/logrus"

	"github.com/Jagrmi-C/gostarted/project/logger"
	"github.com/Jagrmi-C/gostarted/project/models"
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
	
	GetGroupQuery = "SELECT uuid, title FROM groups WHERE uuid=$1"
	GetGroupsQuery = "SELECT * FROM groups"
	CreateGroupQuery = "INSERT INTO groups(title, group_uuid, dt) VALUES($1, $2, $3) RETURNING uuid"
	UpdateGroupQuery = "UPDATE tasks SET title=$1,dt=$2 WHERE uuid=$3"
	DeleteGroupQuery = "DELETE FROM groups WHERE uuid=$1"
)

func init()  {
	logger.LoggerInitialization()
}

func CreateConnection() *pgx.Conn {
	conn, err := pgx.Connect(
		context.Background(),
		os.Getenv("DATABASE_URL"),
	)

	if err != nil {
		// panic(err)
		lr.Error(err)
	}

	// check the connection
	err = conn.Ping(context.Background())

	if err != nil {
		// panic(err)
		lr.Error(err)
	}

	lr.Info("Successfully connected!")
	// return the connection
	return conn
}


// func GetTask(db *pgx.Conn, uuid string, t *models.Task) error {
func GetTask(uuid string, t *models.Task) error {
	conn := CreateConnection()

	defer conn.Close(context.Background())
	err := conn.QueryRow(
		context.Background(),
		GetTaskQuery,
		uuid,
	).Scan(&t.UUID, &t.Title)

	return err
}

func GetTasks() ([]models.Task, error) {
	conn := CreateConnection()

	defer conn.Close(context.Background())
	var tasks []models.Task

	rows, _ := conn.Query(context.Background(), GetTasksQuery)

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

func UpdateTask(t *models.Task) error {
	conn := CreateConnection()

	defer conn.Close(context.Background())

	if (t.Title != "" && t.GroupUUID != "") {
		_, err := conn.Exec(
			context.Background(),
			UpdateTaskQuery,
			t.Title,
			t.GroupUUID,
			t.UUID,
		)
		return err
	} else if (t.Title != "") {
		_, err := conn.Exec(
			context.Background(),
			UpdateTaskQueryIfTitle,
			t.Title,
			t.UUID,
		)
		return err
	} else {
		_, err := conn.Exec(
			context.Background(),
			UpdateTaskQueryIfGroup,
			t.GroupUUID,
			t.UUID,
		)
		return err
	}
}

func CreateTask(t *models.Task) error {
	conn := CreateConnection()

	defer conn.Close(context.Background())

	err := conn.QueryRow(
		context.Background(),
		CreateTaskQuery,
		t.Title,
		t.GroupUUID,
	).Scan(&t.UUID)
	return err
}

func DeleteTask(uuid string) error {
	conn := CreateConnection()

	defer conn.Close(context.Background())

	_, err := conn.Exec(
		context.Background(),
		DeleteTaskQuery,
		uuid,
	)
	return err
}

func GetGroup(uuid string, gr *models.Group) error {
	conn := CreateConnection()

	defer conn.Close(context.Background())

	err := conn.QueryRow(
		context.Background(),
		GetGroupQuery,
		uuid,
	).Scan(&gr.UUID, &gr.Title)

	return err
}

func GetGroups() ([]models.Group, error) {
	conn := CreateConnection()

	defer conn.Close(context.Background())
	var groups []models.Group

	rows, _ := conn.Query(context.Background(), GetGroupsQuery)

	for rows.Next() {
		var group models.Group
		err := rows.Scan(&group.UUID, &group.Title,)
		if err != nil {
			return groups, err
		}
		groups = append(groups, group)
	}

	return groups, rows.Err()
}

func CreateGroup(gr *models.Group) error {
	conn := CreateConnection()

	defer conn.Close(context.Background())

	err := conn.QueryRow(
		context.Background(),
		CreateGroupQuery,
		gr.Title,
		gr.DT,
	).Scan(&gr.UUID)
	return err
}

func UpdateGroup(gr *models.Group) error {
	conn := CreateConnection()

	defer conn.Close(context.Background())

	_, err := conn.Exec(
		context.Background(),
		UpdateGroupQuery,
		gr.Title,
		gr.DT,
		gr.UUID,
	)
	return err
}

func DeleteGroup(uuid string) error {
	conn := CreateConnection()

	defer conn.Close(context.Background())

	_, err := conn.Exec(
		context.Background(),
		DeleteGroupQuery,
		uuid,
	)
	return err
}
