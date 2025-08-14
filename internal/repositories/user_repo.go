package repositories

import(
	"gorm.io/gorm"
	"todo_api/internal/models"
)
type UserRepository struct{
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository{
	return &UserRepository{db:db}
}

func (r *UserRepository) Create(user *models.User) error{
	return r.db.Create(user).Error
}
func (r *UserRepository) GetByID(id uint) (*models.User, error){
	var user models.User
	if err := r.db.First(&user,id).Error; err != nil{
		return nil, err 
	}
	return &user,nil
}
func (r *UserRepository) GetByUsername(username string) (*models.User, error){
	var user models.User
	if err := r.db.Where("username = ?",username).First(&user).Error;err != nil{
		return nil, err
	}
	return &user, nil
}
func (r *UserRepository) Update(user *models.User) error{
	return r.db.Save(user).Error
}
func (r *UserRepository) Delete(id uint) error{
	return r.db.Delete(&models.User{},id).Error
}
