package services

import (
	"time"
	"todo_api/internal/models"
	"todo_api/internal/repositories"
)
type TaskService struct{
	taskRepo *repositories.TaskRepository
}
func NewTaskService(taskRepo *repositories.TaskRepository) *TaskService{
	return &TaskService{taskRepo: taskRepo}
}

func (s *TaskService) CreateTask(userID uint, title, status string, dueDate *time.Time) (*models.Task, error){
	task := &models.Task{
		Title: title,
		UserId: userID,
		Status: status,
		DueDate: dueDate,
		CreatedAt: time.Now(),
	}
	if err := s.taskRepo.Create(task); err!= nil{
		return nil, err
	}
	return task, nil
}
func (s *TaskService) UpdateTask(task *models.Task) error{
	return s.taskRepo.Update(task)
}
func (s *TaskService) DeleteTask(id uint) error{
	return s.taskRepo.Delete(id)
}
func (s *TaskService) GetTaskByID(id uint) (*models.Task, error){
	return s.taskRepo.GetByID(id)
}
func (s *TaskService) GetTasksByUserID(id uint) ([]models.Task, error){
	return s.taskRepo.GetByUserID(id)
}
func (s *TaskService) GetTasksByStatus(userID uint, status string) ([]models.Task, error){
	return s.taskRepo.GetBySatus(status, userID)
} 
func (s *TaskService) GetDueTasks(userID uint, before time.Time) ([]models.Task, error){
	return s.taskRepo.GetDueTasks(userID, before)
}
