package service

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var (
	// ErrUnauthenticated used when there is no autheniticated user in context
	ErrUnauthenticated = errors.New("unauthenticated")
)

//LoginOutput response
type LoginOutput struct {
	Token     string    `json:"token,omitempty"`
	AuthUser  User      `json:"auth_user,omitempty"`
}

//Login provides user login and token generation
func (s *Service) Login(ctx *gin.Context, email string, password string) (LoginOutput, error) {

	var l LoginOutput
	
	query := "SELECT id, username, password, email FROM users WHERE email = $1"
				

	err := s.db.QueryRowContext(ctx, query, email).Scan(&l.AuthUser.ID, &l.AuthUser.Username, &l.AuthUser.Password, &l.AuthUser.Email)

	if err == sql.ErrNoRows {
		return LoginOutput{}, ErrUserNotFound
	}

	if err != nil {
		return LoginOutput{}, fmt.Errorf("could not query select user: %v", err)
	}

	if !comparePasswords(l.AuthUser.Password, []byte(password)) {
		return LoginOutput{}, ErrUserNotFound
	} 

	l.AuthUser.Password = ""
	
	token, err := createToken(l.AuthUser.ID, s.secret)

	if err != nil {
		return l, fmt.Errorf("could not create token for user: %v", err)		
	}

	l.Token = token
	
	return l, nil
} 


func createToken(userid int64, secret string) (string, error) {
  var err error
  //Creating Access Token
  atClaims := jwt.MapClaims{}
  atClaims["authorized"] = true
  atClaims["user_id"] = userid
  atClaims["expiry"] = time.Now().Add(time.Hour * 24).Unix()
  at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
  token, err := at.SignedString([]byte(secret))
  if err != nil {
     return "", err
  }
  return token, nil
}

func comparePasswords(hashedPwd string, plainPwd []byte) bool {
    // Since we'll be getting the hashed password from the DB it
    // will be a string so we'll need to convert it to a byte slice
    byteHash := []byte(hashedPwd)
    err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
    if err != nil {
        return false
    }
    
    return true
}

