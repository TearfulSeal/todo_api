package repositories

import(
	"gorm.io/gorm"
	"todo_api/internal/models"
	"time"
)
type TaskRepository struct{
	db *gorm.DB
}
func NewTaskRepository(db *gorm.DB) *TaskRepository{
	return &TaskRepository{db:db}
}
func (r *TaskRepository) Create(task *models.Task) error{
	return r.db.Create(task).Error 
}
func (r * TaskRepository) GetByID(id uint) (*models.Task, error){
	var task models.Task
	if err := r.db.First(&task,id).Error; err != nil{
		return nil, err
	}
	return &task, nil
}
func (r *TaskRepository) Update(task *models.Task) error{
	return r.db.Save(task).Error 
}
func (r *TaskRepository) Delete(id uint) error{
	return r.db.Delete(&models.Task{},id).Error
}
func (r *TaskRepository) GetByUserID(userId uint) ([]models.Task, error){
	var tasks []models.Task
	query := r.db.Where("user_id = ?",userId)
	if err := query.Find(&tasks).Error;err!=nil{
		return nil, err 
	}
	return tasks,nil 
}
func (r *TaskRepository) GetBySatus(status,userId string) ([]models.Task,error){
	var tasks []models.Task
	query := r.db.Where("status = ? AND user_id = ?",status,userId)
	if err := query.Find(&tasks).Error;err!=nil{
		return nil, err 
	}
	return tasks,nil 
}
func (r *TaskRepository) GetDueTasks(userID uint, before time.Time) ([]models.Task, error) {
    var tasks []models.Task
    if err := r.db.
        Where("user_id = ? AND due_date < ?", userID, before).
        Order("due_date ASC").
        Find(&tasks).Error; err != nil {
        return nil, err
    }
    return tasks, nil
}

