package auth

import (
	"database/sql"
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	DB *sql.DB
}

func (user *UserRepository) InsertUser(data User) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	query := `INSERT INTO 
	users(first_name, username, password) 
	VALUES ($1, $2, $3);`
	stmt, err := user.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(data.FirstName, data.Username, string(bytes))
	if err != nil {
		log.Println(err)
		return err
	}

	i, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if i != 1 {
		return ErrRowsAffected
	}
	log.Println("user inserted success")
	return nil
}

func (user *UserRepository) GetUser(data User) (User, error) {
	var userData User

	query := `SELECT id, username, password
	FROM users WHERE username=$1;`
	err := user.DB.QueryRow(query, data.Username).Scan(
		&userData.ID,
		&userData.Username,
		&userData.Password,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return userData, ErrIncorrectCredentials
		}
		log.Println("Error querying for user:", err)
		return userData, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(data.Password))
	if err != nil {
		return userData, ErrIncorrectCredentials
	}

	return userData, nil
}

func (user *UserRepository) CreatePreference(data User) error {
	var userData User

	query := `SELECT id FROM users WHERE username=$1;`
	err := user.DB.QueryRow(query, data.Username).Scan(&userData.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrObteinData
		}
		log.Println("Error querying user ID:", err)
		return err
	}

	result, err := user.DB.Exec(`INSERT INTO preference(user_id) VALUES ($1);`, userData.ID)
	if err != nil {
		log.Println(err)
		return err
	}

	i, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if i != 1 {
		return errors.New("1 row was expected to be affect")
	}
	log.Println("preference create success")
	return nil
}

func (user *UserRepository) GetPreferenceID(user_id string) (string, error) {
	var preference Preference

	query := `SELECT id
	FROM preference WHERE user_id=$1;`
	err := user.DB.QueryRow(query, user_id).Scan(
		&preference.ID,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", ErrNoPreferencesFound
		}
		log.Println("Error querying for user:", err)
		return "", err
	}
	return preference.ID, nil
}
