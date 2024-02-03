package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	connectDatabase()
	createWebService()
}

func connectDatabase() {
	cfg := getEnvVariable("DB_CONNECTION")

	var err error
	db, err = sql.Open("postgres", cfg)
	if err != nil {
		log.Fatal(err)
	}
	pingDatabase()
	fmt.Println("Connected!")
}

func pingDatabase() {
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
}

func createWebService() {
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine = defineRoutes(engine)
	engine.Run(fmt.Sprintf(":%s", getEnvVariable("PORT")))
}

func defineRoutes(engine *gin.Engine) *gin.Engine {
	engine.GET("/")
	return engine
}

func getEnvVariable(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
