package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"tomaxut/config"
	"tomaxut/server"
	"tomaxut/server/middleware"
	"tomaxut/store/database"
	"tomaxut/store/logger"
	"tomaxut/store/models"
)

// 最终方案-全兼容
func getCurrentAbPath() string {
	dir := getCurrentAbPathByExecutable()
	if strings.Contains(dir, getTmpDir()) {
		return getCurrentAbPathByCaller()
	}
	return dir
}

// 获取系统临时目录，兼容go run
func getTmpDir() string {
	dir := os.Getenv("TEMP")
	if dir == "" {
		dir = os.Getenv("TMP")
	}
	res, _ := filepath.EvalSymlinks(dir)
	return res
}

// 获取当前执行文件绝对路径
func getCurrentAbPathByExecutable() string {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	res, _ := filepath.EvalSymlinks(filepath.Dir(exePath))
	return res
}

// 获取当前执行文件绝对路径（go run）
func getCurrentAbPathByCaller() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}
func main() {
	logger.InitLogger()
	if err := config.Load("config/config.yaml"); err != nil {
		fmt.Println("Failed to load configuration")
		return
	}

	println(getCurrentAbPath())
	logrus.Infoln(getCurrentAbPath())
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
