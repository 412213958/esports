package user

import (
	model "esports/models"
	"esports/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type UserController struct{}

// GetUsers 获取所有用户
func (uc UserController) GetUsers(c *gin.Context) {
	userService := NewUserService()
	users, err := userService.GetAllUsers()
	if err != nil {
		utils.HandleError(c, err, http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, users)
}

// GetUserByID 根据用户 ID 获取用户
func (uc UserController) GetUserByID(c *gin.Context) {
	userID := c.Param("id")
	userService := NewUserService()
	user, err := userService.GetUserByID(userID)
	if err != nil {
		utils.HandleError(c, err, http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, user)
}

// CreateUser 创建新用户
func (uc UserController) CreateUser(c *gin.Context) {
	var user model.User
	if err := c.BindJSON(&user); err != nil {
		utils.HandleError(c, err, http.StatusBadRequest)
		return
	}
	user.CreatedAt = time.Now()
	userService := NewUserService()
	err := userService.CreateUser(&user)
	if err != nil {
		utils.HandleError(c, err, http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusCreated, user)
}

// UpdateUser 更新用户信息
func (uc UserController) UpdateUser(c *gin.Context) {
	userID := c.Param("id")
	var updatedUser model.User
	if err := c.BindJSON(&updatedUser); err != nil {
		utils.HandleError(c, err, http.StatusBadRequest)
		return
	}

	userService := NewUserService()
	err := userService.UpdateUser(userID, &updatedUser)
	if err != nil {
		utils.HandleError(c, err, http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, updatedUser)
}

// DeleteUser 删除用户
func (uc UserController) DeleteUser(c *gin.Context) {
	userID := c.Param("id")
	userService := NewUserService()
	err := userService.DeleteUser(userID)
	if err != nil {
		utils.HandleError(c, err, http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
