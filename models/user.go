package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

type User struct {
	ID      int
	Name    string 
	Age     int
	Address Address
}

var (
	usersDir = "users_saved"
	NUM_OF_FILES = 1
	nextID = 1
)

func AddUser(u User) (User, error) {
	if u.ID != 0 {
		return User{}, errors.New("new user must not include id or it must be set to zero")
	}

	u.ID = nextID
	nextID++

	// Build the file path
	filePath := filepath.Join(usersDir, fmt.Sprintf("%d.txt", u.ID))

	// Marshal the user struct to JSON and handle any errors
	data, err := json.Marshal(u)
	if err != nil {
		return User{}, err
	}

	// Write the JSON data to the file and handle any errors
	err = os.WriteFile(filePath, data, 0644)
	if err != nil {
		return User{}, err
	}

	return u, nil
}

func GetUserByID(id int) (User, error) {

	// Build the file path
	filePath := filepath.Join(usersDir, fmt.Sprintf("%d.txt", id))

	// Read the file
	userJSON, err := os.ReadFile(filePath)
	if err != nil {
		return User{}, err
	}

	// Unmarshal JSON into User struct
	var user User
	if err := json.Unmarshal(userJSON, &user); err != nil {
		return User{}, err
	}

	return user, nil
}
