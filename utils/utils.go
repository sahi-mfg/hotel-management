package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func FindEntityByID[T any](c *gin.Context, db *gorm.DB, idParam string, entity *T) bool {
	if err := db.Where("id = ?", c.Param(idParam)).First(entity).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Resource not found"})
		return false
	}
	return true
}

func GetAllEntities[T any](c *gin.Context, db *gorm.DB, entities *[]T) {
	db.Find(entities)
	c.JSON(http.StatusOK, entities)
}
