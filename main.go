package main

import (
	"to-do-application-api/controller"
	"to-do-application-api/db"
	"to-do-application-api/repository"
	"to-do-application-api/router"
	"to-do-application-api/usecase"
	"to-do-application-api/validator"
)

func main() {
	db := db.NewDB()
	userRepository := repository.NewUserRepository(db)
	taskRepository := repository.NewTaskRepository(db)
	userValidator := validator.NewUserValidator()
	taskValidator := validator.NewTaskValidator()
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	taskUsecase := usecase.NewTaskUsecase(taskRepository, taskValidator)
	userController := controller.NewUserController(userUsecase)
	taskController := controller.NewTaskController(taskUsecase)
	e := router.NewRouter(userController, taskController)
	e.Logger.Fatal(e.Start(":8080"))
}
