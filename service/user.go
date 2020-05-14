package service

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)


var (
	// ErrEmailTaken used when email already exists.
	ErrEmailTaken = errors.New("email is taken")
	// ErrUsernameTaken used when username already exists.
	ErrUsernameTaken = errors.New("username is taken")
	// ErrUserNotFound used when the user not found on the db.
	ErrUserNotFound = errors.New("user not found")
)

// User model
type User struct {
	ID          int64	
	FirstName   string  `db:"first_name" json:"firstName,omitempty" binding:"required,alphanum,min=4,max=20"`
	LastName    string  `db:"last_name" json:"lastName,omitempty" binding:"required,alphanum,min=4,max=20"`
	DateOfBirth int64   `db:"date_of_birth" json:"dateOfBirth,omitempty" validate:"dateOfBirth"` //Date of birth is in epoch
	Address     string  `json:"address,omitempty" binding:"alphanum,min=4,max=50"`
	Email       string  `json:"email,omitempty" binding:"required,email"`
	Username    string  `json:"username,omitempty" binding:"required,alphanum,min=4,max=15"`
	Password    string  `json:"password,omitempty" binding:"required,min=4,max=25"` // crypto password
}

// Users list
func (s *Service) Users(ctx *gin.Context, page int, size int) ([]User, error) {
	log.Println("Users service")

	offset := page * size

	uu := []User{}

	query := "SELECT * FROM users ORDER BY id ASC OFFSET $1 LIMIT $2"

	err := s.db.SelectContext(ctx, &uu, query, offset, size)

	for i := range uu {
		uu[i].Password = ""
	}

	if err != nil {
		return nil, fmt.Errorf("could not list users: %v", err)
	}

	log.Println("List of users", uu)

	return uu, nil
}


// User returns user
func (s *Service) User(ctx *gin.Context, id int) (User, error) {
	log.Println("User service")

	u := User{}

	query := "SELECT * FROM users WHERE id = $1 ORDER BY id ASC"

	err := s.db.GetContext(ctx, &u, query, id)
	
	u.Password = ""	

	if err != nil {
		return User{}, fmt.Errorf("could get user: %v", err)
	}

	log.Println("User returned ", u)

	return u, nil
}


// CreateUser creates user
func (s *Service) CreateUser(ctx *gin.Context, user User) error {

	log.Println("Create User service")
	log.Println("User", user)

	// Hash password
	pass, err := hashAndSalt([]byte(user.Password))

	if err != nil {
		return fmt.Errorf("Could not generate hash from password: %v", err)		
	}
	user.Password = pass

	//Create and execute query
	query := `INSERT INTO users (first_name, last_name, date_of_birth, address, email, username, password) 
	VALUES (:first,:last,:dob,:address,:email,:username,:password)`

	tx := s.db.MustBegin() 

	_, err = tx.NamedExecContext(ctx, query, 
	map[string]interface{}{
		"first": user.FirstName,
		"last": user.LastName,
		"dob":user.DateOfBirth,
		"address":user.Address,
		"email": user.Email,
		"username":user.Username,
		"password":user.Password,
	})

	if err != nil {
		if rb := tx.Rollback(); rb != nil {
			return fmt.Errorf("query failed: %v, unable to abort: %v", err, rb)
		}
		return err
	}
	
	if err := tx.Commit(); err != nil {
		return err
	}

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

func hashAndSalt(pwd []byte) (string, error) {
    
    // Use GenerateFromPassword to hash & salt pwd.
    // MinCost is just an integer constant provided by the bcrypt
    // package along with DefaultCost & MaxCost. 
    // The cost can be any value you want provided it isn't lower
    // than the MinCost (4)
    hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
    if err != nil {
        return "", err
    }
    // GenerateFromPassword returns a byte slice so we need to
    // convert the bytes to a string and return it
    return string(hash), nil
}