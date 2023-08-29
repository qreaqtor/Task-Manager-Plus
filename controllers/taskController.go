package controllers

import (
	"net/http"
	"task-manager-plus/models"
	"task-manager-plus/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskController struct {
	Taskservice services.TaskService
}

func NewTaskController() TaskController {
	return TaskController{
		Taskservice: services.NewTaskService(),
	}
}

func (tc *TaskController) CreateTask(ctx *gin.Context) {
	userId := ctx.MustGet("userId").(primitive.ObjectID)
	var task models.Task
	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	task.UserId = userId
	err := tc.Taskservice.CreateTask(task)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "succes"})
}

// func (ts *TaskController) GetUserTasks(ctx *gin.Context) {
// 	userId := ctx.MustGet("userId").(primitive.ObjectID)

// }

func (tc *TaskController) RegisterTasksRoutes(rg *gin.RouterGroup) {
	rg.POST("/create", tc.CreateTask)
}
