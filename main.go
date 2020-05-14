package main

import (
	"fmt"
	"log"
	"os"

	"github.com/djomlaa/helpee/handler"
	"github.com/djomlaa/helpee/service"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	dbport   = 5433
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
	schema   = "helpee"
)

func main() {

	var (
		port      = env("PORT", "8789")
		origin    = env("ORIGIN", "http://localhost:"+port)
	)
	
	env("SECRET", "totalymeagasecretkey")	

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable search_path=%s",
		host, dbport, user, password, dbname, schema)
	db, err := sqlx.Connect("postgres", psqlInfo)

	if err != nil {
		log.Fatalf("Could not open db connection: %v\n", err)
		return
	}

	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatalf("Could not ping to db: %v\n", err)
		return
	}

	server := gin.New()

	s := service.New(db, origin)

	handler.SetRouter(s, server)

	server.Run(":" + port)

	// TODO graceful shutdown 
}

func env(key, fallbackValue string) string {
	if key == "SECRET" {
		os.Setenv("SECRET", fallbackValue) 
	}
	s := os.Getenv(key)
	if s == "" {
		return fallbackValue
	}

	return s
}