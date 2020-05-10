package service

import (
	"log"
)

// User model
type User struct {
	ID          int64
	FirstName   string `db:"first_name"`
	LastName    string `db:"last_name"`
	DateOfBirth int64  `db:"date_of_birth"`
	Address     string 
	Email       string 
	Username    string 
	Password    string 
}

// Users list
func (s *Service) Users() ([]User, error) {
	log.Println("Users service")

	uu := []User{}

	query := "SELECT * FROM users ORDER BY id ASC"

	err := s.db.Select(&uu, query)

	if err != nil {
		return nil, err
	}

	log.Println("List of users", uu)

	return uu, nil
}
