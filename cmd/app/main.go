package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"tomaxut/config"
	"tomaxut/server"
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
	)

	gin.SetMode(config.Get().GinMode)

	router := gin.Default()

	server.InitRouter(router)

	_ = router.Run(config.Get().Addr)
}
