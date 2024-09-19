package database

import (
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib" // load pgx driver for PostgreSQL
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)




func OpenDBConnection() (*sqlx.DB, error){
	databaseUrl := connectionURLBuilder() 

	db, err := sqlx.Connect("pgx", databaseUrl)

	if err != nil {
		fmt.Printf("Could not connect to the database: %s", err);
		return nil, err
	}

	return db, nil

}


func connectionURLBuilder() string {
	godotenv.Load()
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s",
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_NAME"),
	)
}