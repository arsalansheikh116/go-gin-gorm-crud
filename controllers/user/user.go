package controllers

import (
	"net/http"
	"strings"

	"go-crud/config"
	models "go-crud/models/user"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Handlers

type IndexedUser struct {
	Index int         `json:"index"`
	User  models.User `json:"user"`
}

func GetUsers(c *gin.Context) {
	var users []models.User
	if err := config.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Create a slice of IndexedUser
	indexedUsers := make([]IndexedUser, len(users))
	for i, user := range users {
		indexedUsers[i] = IndexedUser{
			Index: i + 1,
			User:  user,
		}
	}

	c.JSON(http.StatusOK, indexedUsers)
}

func GetUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	// Bind the JSON request body to a User struct
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash the password before saving
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not hash password"})
		return
	}
	// Store the hashed password in the user object
	user.Password = string(hashedPassword)

	// Create the user in the database
	if err := config.DB.Create(&user).Error; err != nil {
		// Check if the error is due to duplicate name or email
		if strings.Contains(err.Error(), "users_name_key") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Name already exists"})
			return
		} else if strings.Contains(err.Error(), "users_email_key") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
			return
		}
		// For other errors, return internal server error
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return success response
	response := gin.H{
		"data": gin.H{
			"data": gin.H{
				"id":    user.ID,
				"name":  user.Name,
				"email": user.Email,
			},
			"message": "User created successfully",
		},
	}
	c.JSON(http.StatusCreated, response)
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&user)
	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.User{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
