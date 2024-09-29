package user

import (
	dbConetion "api_pattern_go/api/database/conection"
	"api_pattern_go/api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	dbConetion.DB.Create(&user)
	c.JSON(http.StatusCreated, user)
}

func GetUsers(c *gin.Context) {
	var users []models.User
	dbConetion.DB.Find(&users)
	c.JSON(http.StatusOK, users)
}
