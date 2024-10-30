package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func PGConnection() (*sql.DB, error) {
	DB_HOST := os.Getenv("PG_HOST")
	DB_USER := os.Getenv("PG_USER")
	DB_PASSWORD := os.Getenv("PG_PASSWORD")
	DB_NAME := os.Getenv("PG_NAME")
	DB_PORT := os.Getenv("PG_PORT")

	uri := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		DB_USER,
		DB_PASSWORD,
		DB_HOST,
		DB_PORT,
		DB_NAME,
	)
	db, err := sql.Open("postgres", uri)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func PGMigrate(db *sql.DB) error {
	const create = `
	BEGIN;

	CREATE TABLE IF NOT EXISTS users (
		id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
		first_name VARCHAR(100),
		username VARCHAR(50) UNIQUE,
		password VARCHAR(100),
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS preference (
		id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
		user_id uuid REFERENCES users(id) ON DELETE CASCADE,
		UNIQUE(user_id)
	);

	CREATE TABLE IF NOT EXISTS genre_score (
		id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
		preference_id uuid REFERENCES preference(id) ON DELETE CASCADE,
		name VARCHAR,
		score DECIMAL(5, 2),
		UNIQUE(preference_id, name)
	);

	CREATE TABLE IF NOT EXISTS protagonist_score (
		id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
		preference_id uuid REFERENCES preference(id) ON DELETE CASCADE,
		name VARCHAR,
		score DECIMAL(5, 2),
		UNIQUE(preference_id, name)
	);

	CREATE TABLE IF NOT EXISTS director_score (
		id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
		preference_id uuid REFERENCES preference(id) ON DELETE CASCADE,
		name VARCHAR,
		score DECIMAL(5, 2),
		UNIQUE(preference_id, name)
	);

	CREATE TABLE IF NOT EXISTS history (
		user_id uuid REFERENCES users(id) ON DELETE CASCADE,
		movie_id VARCHAR,
		watched_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		PRIMARY KEY (user_id, movie_id)
	);

	COMMIT;`

	result, err := db.Exec(create)
	if err != nil {
		return err
	}
	_, err = result.RowsAffected()
	if err != nil {
		return err
	}

	log.Println("\033[32mSUCCESS TO CREATE TABLES\033[0m")

	return nil
}
