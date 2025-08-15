package handlers

import (
	"todo_api/internal/services"
	"todo_api/pkg/jwt"

	"github.com/gin-gonic/gin"
)


type AuthHandler struct {
	userService *services.UserService 
}

func NewAuthHandler(userService *services.UserService) *AuthHandler{
	return &AuthHandler{userService: userService}
}
type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *AuthHandler) Register(c *gin.Context){
	var req RegisterRequest
	if err := c.BindJSON(&req); err != nil{
		c.JSON(400, gin.H{"error":err.Error()})
		return
	}
	user, err := h.userService.CreateUser(req.Username,req.Password)
	if err != nil {
		c.JSON(500, gin.H{"error":"cannot create user"})
		return
	}
	token, err := jwt.GenerateToken(user.ID)
	if err != nil {
		c.JSON(500, gin.H{"error":"cannot generate token"})
		return
	}
	c.JSON(200, gin.H{"token":token})
}
