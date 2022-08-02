package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"maxotm/config"
	"maxotm/server"
	"maxotm/store/database"
	"maxotm/store/logger"
	"maxotm/store/models"
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
		&models.AppUpgrade{},
	)

	gin.SetMode(config.Get().GinMode)

	router := gin.Default()

	server.InitRouter(router)
	/*if err := endless.ListenAndServe(config.Get().Addr, router); err != nil {
		fmt.Println(err.Error())
	}*/

	_ = router.Run(config.Get().Addr)
}
