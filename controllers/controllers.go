package controllers

import (
	"hotel-management/database"
	"hotel-management/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Welcome(c *gin.Context) {
	greeting := "Welcome on our hotel management website."
	c.JSON(http.StatusOK, greeting)
}

// Afficher toutes les chambres
func GetChambres(c *gin.Context) {
	var chambres []models.Chambre
	database.DB.Find(&chambres)
	c.JSON(http.StatusOK, chambres)
}

// Ajouter une chambre
func CreateChambre(c *gin.Context) {
	var chambre models.Chambre
	if err := c.ShouldBindJSON(&chambre); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&chambre)
	c.JSON(http.StatusOK, chambre)
}

// Afficher les clients
func GetClients(c *gin.Context) {
	var clients []models.Client
	database.DB.Find(&clients)
	c.JSON(http.StatusOK, clients)
}

// Ajouter un nouveau client
func NewClient(c *gin.Context) {
	var client models.Client
	if err := c.ShouldBindJSON(&client); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&client)
	c.JSON(http.StatusOK, client)
}

// Supprimer les client qui ont checkout
func DeleteClient(c *gin.Context) {

}

func UpdateClient(c *gin.Context) {

}

// Afficher les reservations
func GetRservations(c *gin.Context) {
	var reservation models.Reservation
	database.DB.Find(&reservation)
	c.JSON(http.StatusOK, reservation)
}

// Nouvelles Reservations
func NewReservation(c *gin.Context) {
	var reservation models.Reservation
	if err := c.ShouldBindJSON(&reservation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	database.DB.Create(&reservation)
}

func UpdateRservation(c *gin.Context) {

}

func DeleteReservation(c *gin.Context) {

}

// Afficher les status d'occupation des chambres

func Occupation(c *gin.Context) {

}
