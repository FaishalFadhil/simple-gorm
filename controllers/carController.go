package controllers

import (
	"fmt"
	"gin-api/database"
	"gin-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllCars(c *gin.Context) {
	var db = database.GetDB()

	var cars []models.Car
	err := db.Preload("Owner").Find(&cars).Error

	if err != nil {
		fmt.Println("Error getting car datas :", err.Error())
	}

	c.JSON(http.StatusOK, gin.H{"data": cars})
}

func GetOneCars(c *gin.Context) {
	var db = database.GetDB()

	var carOne models.Car
	// err := db.Table("Car").Where("Id = ?", c.Param("id")).First(&car).Error
	err := db.First(&carOne, "Id = ?", c.Param("id")).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data One": carOne})
}

func CreateCars(c *gin.Context) {
	var db = database.GetDB()
	// Validate input
	var input models.Car
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	carinput := models.Car{Brand: input.Brand, Price: input.Price, Type: input.Type, Owner: input.Owner}
	db.Create(&carinput)

	c.JSON(http.StatusOK, gin.H{"data": carinput})
}

func UpdateCars(c *gin.Context) {
	var db = database.GetDB()

	var car models.Car
	err := db.First(&car, "Id = ?", c.Param("id")).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	// Validate input
	var input models.Car
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Model(&car).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": car})
}

func DeleteCars(c *gin.Context) {
	var db = database.GetDB()
	// Get model if exist
	var carDelete models.Car
	err := db.First(&carDelete, "Id = ?", c.Param("id")).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&carDelete)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

func GetAllUsers(c *gin.Context) {
	var db = database.GetDB()

	var users []models.User
	err := db.Find(&users).Error

	if err != nil {
		fmt.Println("Error getting user datas :", err.Error())
	}

	c.JSON(http.StatusOK, gin.H{"data": users})
}

func GetOneUser(c *gin.Context) {
	var db = database.GetDB()

	var user models.User
	// err := db.Table("Car").Where("Id = ?", c.Param("id")).First(&car).Error
	err := db.First(&user, "Id = ?", c.Param("id")).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data One": user})
}

func CreateUser(c *gin.Context) {
	var db = database.GetDB()
	// Validate input
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	userInput := models.User{Name: input.Name, Address: input.Address, UserId: input.UserId}
	db.Create(&userInput)

	c.JSON(http.StatusOK, gin.H{"data": userInput})
}

func UpdateUser(c *gin.Context) {
	var db = database.GetDB()

	var user models.User
	err := db.First(&user, "UserId = ?", c.Param("id")).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	// Validate input
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Model(&user).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func DeleteUser(c *gin.Context) {
	var db = database.GetDB()
	// Get model if exist
	var userDelete models.User
	err := db.First(&userDelete, "UserId = ?", c.Param("id")).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&userDelete)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
