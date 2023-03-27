package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"tomaxut/config"
	"tomaxut/server"
	"tomaxut/server/middleware"
	"tomaxut/store/database"
	"tomaxut/store/logger"
	"tomaxut/store/models"
)

func main() {
	logger.InitLogger()
	if err := config.Load("config/config.yaml"); err != nil {
		fmt.Println("Failed to load configuration")
		return
	}

	db, err := database.InitDB()
	if err != nil {
		fmt.Println("err open databases")
		return
	}

	_ = db.AutoMigrate(
		&models.User{},
		&models.Post{},
	)

	gin.SetMode(config.Get().GinMode)

	router := gin.Default()
	router.Use(middleware.Cors())
	server.InitRouter(router)

	_ = router.Run(config.Get().Addr)
}
