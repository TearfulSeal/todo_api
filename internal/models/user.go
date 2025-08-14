package models
import"time"
type User struct{
	ID uint `gorm:"primaryKey"`
	Username string `gorm:"uniqueIndex;size:50"`
	Password string
	CreatedAt time.Time
}
