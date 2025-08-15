package handlers

import (
	"todo_api/internal/services"
	"strconv"
	"github.com/gin-gonic/gin"
)


type UserHandler struct{
	userService *services.UserService
}
func NewUserHandler(userService *services.UserService) *UserHandler{
	return &UserHandler{userService: userService}
}
func (h *UserHandler) DeleteUser(c *gin.Context){
	tokenUserID, exists := c.Get("userID")
	if !exists {
		c.JSON(401,gin.H{"error":"unauthorized"})
		return 
	}
	userID := tokenUserID.(uint)
	if err := h.userService.DeleteUser(userID);err !=nil{
		c.JSON(500,gin.H{"error":"failed to delete user"})
		return 
	}
	c.JSON(200,gin.H{"message":"user deleted"})

}
func (h *UserHandler) UpdateUser(c *gin.Context){
	tokenUserID, exists := c.Get("userID")
	if !exists {
		c.JSON(401,gin.H{"error":"unauthorized"})
		return 
	}
	userID := tokenUserID.(uint)
	var input struct{
		Username *string `json:"username"`
		Password *string `json:"password"`
	}
	if err := c.ShouldBindJSON(&input);err!=nil{
		c.JSON(400,gin.H{"error":"invalid input"})
		return 
	}
	user, err := h.userService.GetUserByID(userID)
	if err!=nil{
		c.JSON(404, gin.H{"error":"user not found"})
		return 
	}
	if input.Username!=nil{
		user.Username=*input.Username
	}
	if input.Password!=nil{
		user.Password=*input.Password
	}
	if err := h.userService.UpdateUser(user); err!=nil{
		c.JSON(500,gin.H{"error":"failed to update user"})
		return 
	}
	c.JSON(200,user)
}
func (h *UserHandler) GetUserByID(c *gin.Context){
	userIDParam := c.Param("id")
	userID, _ := strconv.ParseUint(userIDParam,10,64)
	user, err := h.userService.GetUserByID(uint(userID))
	if err!=nil{
		c.JSON(404,gin.H{"error":"user not found"})
		return 
	}
	c.JSON(200,user)
}
func (h *UserHandler) GetUserByUsername(c *gin.Context){
	username := c.Param("username")
	user, err := h.userService.GetUserByUsername(username)
	if err!=nil{
		c.JSON(404,gin.H{"error":"user not found"})
		return 
	}
	c.JSON(200,user)
}
