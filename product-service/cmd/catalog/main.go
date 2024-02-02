package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func main() {
	sql.Open("postgres", "host=lo")

	engine := gin.New()
	engine.Use(gin.Recovery())

	engine.Run(":8080")
}
