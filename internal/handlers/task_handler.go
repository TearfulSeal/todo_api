package handlers

import (
	"strconv"
	"time"
	"todo_api/internal/services"

	"github.com/gin-gonic/gin"
)


type TaskHandler struct{
	taskService *services.TaskService
}
func NewTaskHandler(taskService *services.TaskService) *TaskHandler{
	return &TaskHandler{taskService: taskService}
}
func (h *TaskHandler) GetTasks(c *gin.Context){
	tokenUserID, exists := c.Get("userID")
	if !exists {
		c.JSON(401,gin.H{"error":"unauthorized"})
		return 
	}
	userID := tokenUserID.(uint)
	tasks,err := h.taskService.GetTasksByUserID(userID)
	if err != nil{
		c.JSON(500, gin.H{"error":"failed to get tasks"})
		return 
	}
	c.JSON(200,tasks)

}
func (h *TaskHandler) GetTasksByID(c *gin.Context){
	tokenUserID, exists := c.Get("userID")
	if !exists {
		c.JSON(401,gin.H{"error":"unauthorized"})
		return 
	}
	userID := tokenUserID.(uint)
	taskIDParam := c.Param("id")
	taskID, _ := strconv.ParseUint(taskIDParam,10,64)
	task, err := h.taskService.GetTaskByID(uint(taskID))
	if err!= nil{
		c.JSON(404,gin.H{"error":"task not found"})
		return 
	}
	if task.UserId != userID{
		c.JSON(403, gin.H{"error":"user does not have access"})
		return 
	}
	c.JSON(200,task)
	
}
func (h *TaskHandler) GetTasksByStatus(c *gin.Context){
	tokenUserID, exists := c.Get("userID")
	if !exists {
		c.JSON(401,gin.H{"error":"unauthorized"})
		return 
	}
	userID := tokenUserID.(uint)
	status := c.Param("status")
	tasks, err := h.taskService.GetTasksByStatus(userID,status)
	if err != nil{
		c.JSON(404,gin.H{"error":"tasks not found"})
		return 
	}
	c.JSON(200,tasks)
}
func (h *TaskHandler) CreateTask(c *gin.Context){
	tokenUserID, exists := c.Get("userID")
	if !exists {
		c.JSON(401,gin.H{"error":"unauthorized"})
		return 
	}
	userID := tokenUserID.(uint)
	var input struct{
		Title string `json:"title" binding:"required"`
		Status string `json:"status" binding:"required"`
		DueDate *time.Time `json:"due_date"`
	}	
	if err := c.ShouldBindJSON(&input); err != nil{
		c.JSON(400,gin.H{"error":"invalid input"})
		return 
	}
	task, err := h.taskService.CreateTask(userID, input.Title,input.Status,input.DueDate)
	if err != nil{
		c.JSON(500,gin.H{"error":"failed to create task"})
		return 
	}
	c.JSON(201,task)
}

func (h *TaskHandler) UpdateTask(c *gin.Context){
	tokenUserID, exists := c.Get("userID")
	if !exists {
		c.JSON(401,gin.H{"error":"unauthorized"})
		return 
	}
	userID := tokenUserID.(uint)
	taskIDParam := c.Param("id")
	taskID, _ := strconv.ParseUint(taskIDParam,10,64)
	task, err := h.taskService.GetTaskByID(uint(taskID))
	if err != nil{
		c.JSON(404,gin.H{"error":"task not found"})
		return
	}
	if task.UserId != userID{
		c.JSON(403, gin.H{"error":"user does not have access"})
		return 
	}

	var input struct{
		Title *string `json:"title"`
		Status *string `json:"status"`
		DueDate *time.Time `json:"due_date"`
	}
	if err := c.ShouldBindJSON(&input); err!=nil{
		c.JSON(400,gin.H{"error":"invalid input"})
		return 
	}
	if input.DueDate != nil{
		task.DueDate=input.DueDate
	}
	if input.Status!=nil{
		task.Status=*input.Status
	}
	if input.Title!=nil{
		task.Title=*input.Title
	}
	if err:= h.taskService.UpdateTask(task);err!=nil{
		c.JSON(500,gin.H{"error":"failed to update task"})
		return 
	}
	c.JSON(200,task)

}
func (h *TaskHandler) DeleteTask(c *gin.Context){
	tokenUserID, exists := c.Get("userID")
	if !exists {
		c.JSON(401,gin.H{"error":"unauthorized"})
		return 
	}
	userID := tokenUserID.(uint)
	taskIDParam := c.Param("id")
	taskID, _ := strconv.ParseUint(taskIDParam,10,64)
	task, err := h.taskService.GetTaskByID(uint(taskID))
	if err != nil{
		c.JSON(404,gin.H{"error":"task not found"})
		return
	}
	if task.UserId != userID{
		c.JSON(403, gin.H{"error":"user does not have access"})
		return 
	}

	if err:= h.taskService.DeleteTask(task.ID);err!=nil{
		c.JSON(500,gin.H{"error":"failed to delete task"})
		return 
	}
	c.JSON(200,task)

}
