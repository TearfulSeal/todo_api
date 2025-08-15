package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	//	"github.com/gin-gonic/gin"
	"fmt"
	"log"
	"time"
	"todo_api/config"
	"todo_api/internal/models"
	"todo_api/internal/repositories"
	"todo_api/pkg/jwt"
)


func main(){
	cfg := config.LoadConfig()
	jwt.Init(cfg.JWTSecret,cfg.JWTExpirationHours)
//	r := gin.Default()
	db, err := gorm.Open(sqlite.Open(cfg.DBPath),&gorm.Config{})
	if err != nil{
		log.Fatalf("failed to connect db %v",err)
	}

	err = db.AutoMigrate(&models.Task{},&models.User{})
	if err != nil{
		log.Fatalf("automigration failed %v",err)
	}
	userRepo := repositories.NewUserRepository(db)
	taskRepo := repositories.NewTaskRepository(db)

	user := &models.User{
		Username:"admin1",
		Password:"123123",
		CreatedAt:time.Now(),
	}
	if err := userRepo.Create(user); err!=nil{
		log.Fatalf("failed create user %v",err)
	}
	due := time.Now().Add(24*time.Hour)
	task := &models.Task{
		Title:"code api",
		Status:"pending",
		DueDate: &due,
		UserId:user.ID,
		CreatedAt:time.Now(),
	}
	if err:= taskRepo.Create(task);err!=nil{
		log.Fatalf("failed create task %v",err)
	}
	tasks, err := taskRepo.GetByUserID(user.ID)
	if err != nil{
		log.Fatalf("failed get by user id %v",err)
	}
	fmt.Println("tasks ",tasks)



//	r.Run(":"+cfg.ServerPort)
}
