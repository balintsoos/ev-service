package db

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

var DB *pgx.Conn

func InitDatabase() {
	dsn := "postgres://postgres:password@localhost:5432/ev-service"
	var err error
	DB, err = pgx.Connect(context.Background(), dsn)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	_, err = DB.Exec(context.Background(), `
        CREATE TABLE IF NOT EXISTS electric_vehicles (
            id SERIAL PRIMARY KEY,
            make VARCHAR(50),
            model VARCHAR(50),
            year INT,
            battery_capacity INT,
            range_km INT,
            price NUMERIC
        )
    `)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}

	fmt.Println("Database connection and table setup successful!")
}

func CloseDatabase(ctx context.Context) {
	DB.Close(ctx)
}
