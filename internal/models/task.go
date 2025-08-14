package models
import"time"
type Task struct{
	ID uint `gorm:"primaryKey" json:"id"`
	Title string `gorm:"size:255;not null" json:"title"`
	Status string `json:"status"`
	DueDate *time.Time `json:"due_date,omitempty"`
	UserId uint `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

