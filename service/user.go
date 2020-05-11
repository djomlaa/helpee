package service

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)


var (
	// ErrEmailTaken used when email already exists.
	ErrEmailTaken = errors.New("email is taken")
	// ErrUsernameTaken used when username already exists.
	ErrUsernameTaken = errors.New("username is taken")
)

// User model
type User struct {
	ID          int64
	FirstName   string `db:"first_name" json:"firstName"`
	LastName    string `db:"last_name" json:"lastName"`
	DateOfBirth int64  `db:"date_of_birth" json:"dateOfBirth"`
	Address     string  `json:"address"`
	Email       string  `json:"email"`
	Username    string  `json:"username"`
	Password    string  `json:"password"`
}

// Users list
func (s *Service) Users() ([]User, error) {
	log.Println("Users service")

	uu := []User{}

	query := "SELECT * FROM users ORDER BY id ASC"

	err := s.db.Select(&uu, query)

	if err != nil {
		return nil, fmt.Errorf("could not list users: %v", err)
	}

	log.Println("List of users", uu)

	return uu, nil
}


// CreateUser creates user
func (s *Service) CreateUser(ctx *gin.Context, user User) error {

	log.Println("Create User service")
	log.Println("User", )

	query := `INSERT INTO users (first_name, last_name, date_of_birth, address, email, username, password) 
	VALUES (:first,:last,:dob,:address,:email,:username,:password)`

	_, err := s.db.NamedExec(query, 
	map[string]interface{}{
		"first": user.FirstName,
		"last": user.LastName,
		"dob":user.DateOfBirth,
		"address":user.Address,
		"email": user.Email,
		"username":user.Username,
		"password":user.Password,
	})

	unique := isUniqueViolation(err)

	if unique && strings.Contains((err.Error()), "email") {
		return ErrEmailTaken
	}

	if unique && strings.Contains((err.Error()), "username") {
		return ErrUsernameTaken
	}

	if err != nil {
		return fmt.Errorf("could not insert user: %v", err)
	}

	return nil
}
