package db

import (
	"context"
	"fmt"
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
	GetTasksQueryByGroup   = "SELECT uuid, title FROM tasks where group_uuid=$1"
	UpdateTaskQuery = "UPDATE tasks SET title=$1,group_uuid=$2 WHERE uuid=$3"
	DeleteTaskQuery = "DELETE FROM tasks WHERE uuid=$1"
	CreateTaskQuery = "INSERT INTO tasks(title, group_uuid) VALUES($1, $2) RETURNING uuid"
	
	GetGroupQuery = "SELECT uuid, title FROM groups WHERE uuid=$1"
	GetGroupsQuery = "SELECT uuid, title FROM groups"
	CreateGroupQuery = "INSERT INTO groups(title, dt) VALUES($1, $2) RETURNING uuid"
	UpdateGroupQuery = "UPDATE groups SET title=$1,dt=$2 WHERE uuid=$3"
	DeleteGroupQuery = "DELETE FROM groups WHERE uuid=$1"

	GetTimeFrameQuery = "SELECT * FROM timeframes WHERE uuid=$1"
	CreateTimeFrameQuery = "INSERT INTO timeframes(task_uuid, dtfrom, dtto) VALUES($1, $2, $3) RETURNING uuid"
	DeleteTimeFrameQuery = "DELETE FROM timeframes WHERE uuid=$1"

	GetTimeFramesQuery = "SELECT dtfrom, dtto FROM timeframes WHERE task_uuid=$1"
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
		lr.Error(err)
		panic(err)
	}

	// check the connection
	err = conn.Ping(context.Background())

	if err != nil {
		lr.Error(err)
		panic(err)
	}

	lr.Debug("Successfully connected!")
	// return the connection
	return conn
}

func GetTasks(uuid string, t *[]models.Task) error {
	conn := CreateConnection()

	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(), GetTasksQueryByGroup, uuid)

	if err != nil {
		return err
	}

	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.UUID, &task.Title,)
		if err != nil {
			return err
		}

		*t = append(*t, task)
	}

	return err
}

func GetTaskInfo(uuid string, t *models.TaskInformation) error {
	conn := CreateConnection()

	defer conn.Close(context.Background())
	err := conn.QueryRow(
		context.Background(),
		GetTaskQuery,
		uuid,
	).Scan(&t.UUID, &t.Title)

	if err != nil {
		return err
	}

	var ttms []models.TaskTimeFrame
	err = GetTimeFramesByTask(t.UUID, &ttms)
	if err != nil {
		return err
	}

	return err
}

func GetTasksInfo(t *[]models.TaskInformation) error {
	conn := CreateConnection()

	defer conn.Close(context.Background())

	rows, _ := conn.Query(context.Background(), GetTasksQuery)

	for rows.Next() {
		var task models.TaskInformation
		err := rows.Scan(&task.UUID, &task.Title,)
		if err != nil {
			return err
		}

		var ttms []models.TaskTimeFrame
		err = GetTimeFramesByTask(task.UUID, &ttms)
		task.TimeFrames = ttms
		if err != nil {
			return err
		}

		*t = append(*t, task)
	}

	return rows.Err()
}

func UpdateTask(t *models.Task) error {
	conn := CreateConnection()

	defer conn.Close(context.Background())

	_, err := conn.Exec(
		context.Background(),
		UpdateTaskQuery,
		t.Title,
		t.GroupUUID,
		t.UUID,
	)
	return err
}

func CreateTask(t *models.Task) error {
	conn := CreateConnection()
	defer conn.Close(context.Background())

	tx, err := conn.Begin(context.Background())
	if err != nil {
		lr.Error(err)
		return err
	}

	defer func() {
        if err != nil {
			_ = tx.Rollback(context.Background())
            return
        }
        err = tx.Commit(context.Background())
    }()

	err = tx.QueryRow(
		context.Background(),
		CreateTaskQuery,
		t.Title,
		t.GroupUUID,
	).Scan(&t.UUID)

	if err != nil {
		lr.Error(err)
		return err
	}
	return err
}

func DeleteTask(uuid string) error {
	conn := CreateConnection()

	defer conn.Close(context.Background())

	tx, err := conn.Begin(context.Background())
	if err != nil {
		lr.Error(err)
		return err
	}

	defer func() {
        if err != nil {
			_ = tx.Rollback(context.Background())
            return
        }
        err = tx.Commit(context.Background())
    }()

	_, err = tx.Exec(
		context.Background(),
		DeleteTaskQuery,
		uuid,
	)
	return err
}

func GetGroup(uuid string, gr *models.GroupInformation) error {
	conn := CreateConnection()

	defer conn.Close(context.Background())

	err := conn.QueryRow(
		context.Background(),
		GetGroupQuery,
		uuid,
	).Scan(&gr.UUID, &gr.Title)

	if err != nil {
		fmt.Println("1")
		return err
	}

	var tasks []models.Task
	err = GetTasks(uuid, &tasks)
	if err != nil {
		fmt.Println("12")
		return err
	}
	gr.Tasks = tasks

	return err
}

func GetGroups(gr *[]models.GroupInformation) error {
	conn := CreateConnection()

	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(), GetGroupsQuery)

	if err != nil {
		return err
	}

	for rows.Next() {
		var group models.GroupInformation
		err := rows.Scan(&group.UUID, &group.Title,)
		if err != nil {
			return err
		}
		var tasks []models.Task
		err = GetTasks(group.UUID, &tasks)
		if err != nil {
			return err
		}

		group.Tasks = tasks
		*gr = append(*gr, group)
	}

	return rows.Err()
}

func CreateGroup(gr *models.Group) error {
	conn := CreateConnection()

	defer conn.Close(context.Background())

	tx, err := conn.Begin(context.Background())
	if err != nil {
		lr.Error(err)
		return err
	}

	defer func() {
        if err != nil {
			_ = tx.Rollback(context.Background())
            return
        }
        err = tx.Commit(context.Background())
    }()

	err = tx.QueryRow(
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

	tx, err := conn.Begin(context.Background())
	if err != nil {
		lr.Error(err)
		return err
	}

	defer func() {
        if err != nil {
			_ = tx.Rollback(context.Background())
            return
        }
        err = tx.Commit(context.Background())
    }()

	_, err = tx.Exec(
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

	tx, err := conn.Begin(context.Background())
	if err != nil {
		lr.Error(err)
		return err
	}

	defer func() {
        if err != nil {
			_ = tx.Rollback(context.Background())
            return
        }
        err = tx.Commit(context.Background())
    }()

	_, err = tx.Exec(
		context.Background(),
		DeleteGroupQuery,
		uuid,
	)
	return err
}

func GetTimeFrame(uuid string, tf *models.TimeFrame) error {
	conn := CreateConnection()

	defer conn.Close(context.Background())

	err := conn.QueryRow(
		context.Background(),
		GetTimeFrameQuery,
		uuid,
	).Scan(&tf.UUID, &tf.FROM, &tf.TO)

	return err
}

func CreateTimeFrame(tf *models.TimeFrame) error {
	conn := CreateConnection()

	defer conn.Close(context.Background())

	tx, err := conn.Begin(context.Background())
	if err != nil {
		lr.Error(err)
		return err
	}

	defer func() {
        if err != nil {
			_ = tx.Rollback(context.Background())
            return
        }
        err = tx.Commit(context.Background())
    }()

	err = tx.QueryRow(
		context.Background(),
		CreateTimeFrameQuery,
		tf.TaskUUID,
		tf.FROM,
		tf.TO,
	).Scan(&tf.UUID)
	return err
}

func DeleteTimeFrame(uuid string) error {
	conn := CreateConnection()

	defer conn.Close(context.Background())

	tx, err := conn.Begin(context.Background())
	if err != nil {
		lr.Error(err)
		return err
	}

	defer func() {
        if err != nil {
			_ = tx.Rollback(context.Background())
            return
        }
        err = tx.Commit(context.Background())
    }()

	_, err = tx.Exec(
		context.Background(),
		DeleteTimeFrameQuery,
		uuid,
	)
	return err
}

func GetTimeFramesByTask(taskUUID string, ttms *[]models.TaskTimeFrame) error {
	conn := CreateConnection()

	defer conn.Close(context.Background())

	rows, err := conn.Query(
		context.Background(),
		GetTimeFramesQuery,
		taskUUID,
	)
	if err != nil {
		lr.Error(err)
		return err
	}

	for rows.Next() {
		var ttm models.TaskTimeFrame
		err := rows.Scan(&ttm.FROM, &ttm.TO,)
		if err != nil {
			lr.Error(err)
			return err
		}
		*ttms = append(*ttms, ttm)
	}

	return rows.Err()
}

func GetGoTasks(output chan []models.TaskInformation) {
	conn := CreateConnection()

	defer conn.Close(context.Background())
	var tasks []models.TaskInformation

	rows, _ := conn.Query(context.Background(), GetTasksQuery)

	for rows.Next() {
		var task models.TaskInformation
		err := rows.Scan(&task.UUID, &task.Title,)
		if err != nil {
			fmt.Println(err)
		}
		var ttms []models.TaskTimeFrame
		err = GetTimeFramesByTask(task.UUID, &ttms)
		if err != nil {
			fmt.Println(err)
		}
		tasks = append(tasks, task)
	}

	output <- tasks
}

func GetGoTask(uuid string, t *models.Task) error {
	conn := CreateConnection()

	defer conn.Close(context.Background())
	err := conn.QueryRow(
		context.Background(),
		GetTaskQuery,
		uuid,
	).Scan(&t.UUID, &t.Title)

	return err
}
