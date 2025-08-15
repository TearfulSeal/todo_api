package services

import(
	"todo_api/internal/models"
	"todo_api/internal/repositories"
	"time"
	"errors"
)
type UserService struct{
	userRepo *repositories.UserRepository
}

func NewUserService(userRepo *repositories.UserRepository) *UserService{
	return &UserService{userRepo: userRepo}
}

func (s *UserService) CreateUser(username, password string) (*models.User, error){
	if username == "" || password == ""{
		return nil, errors.New("username and password required")
	}
	user := &models.User{
		Username: username,
		Password: password,
		CreatedAt: time.Now(),
	}
	if err := s.userRepo.Create(user); err!= nil{
		return nil, err
	}
	return user, nil
}
func (s *UserService) UpdateUser(user *models.User) error{
	return s.userRepo.Update(user)
}
func (s *UserService) DeleteUser(id uint) error{
	return s.userRepo.Delete(id)
}
func (s *UserService) GetUserByID(id uint) (*models.User, error){
	return s.userRepo.GetByID(id)
}
func (s *UserService) GetUserByUsername(username string) (*models.User,error){
	return s.userRepo.GetByUsername(username)
}
