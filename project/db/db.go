package db

import (
	"fmt"
	"os"
	"context"
	"github.com/jackc/pgx/v4"
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