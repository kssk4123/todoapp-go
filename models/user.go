package models

import (
	"database/sql"
	"errors"
        "log"
	"golang.org/x/crypto/bcrypt"
)

// User represents a user in the system.
type User struct {
	ID       int
	Username string
	Password string
}

// CreateUser inserts a new user into the database.
func CreateUser(db *sql.DB, username, password string) (int, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
                log.Println(err)
		return 0,  err
	}

        result, err := db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", username, hashedPassword)
        if err != nil {
                return 0, err
        }

        // 挿入したユーザーの id を取得
        userID, err := result.LastInsertId()
        log.Println(userID)
        if err != nil {
                return 0, err
        }
	return int(userID), err
}

// AuthenticateUser checks username and password.
func AuthenticateUser(db *sql.DB, username, password string) (bool, error) {
	var passwordHash string
	err := db.QueryRow("SELECT password FROM users WHERE username = ?", username).Scan(&passwordHash)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, errors.New("user not found")
		}
		return false, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password)); err != nil {
		return false, errors.New("invalid password")
	}

	return true, nil
}
