package controllers

import (
	"hotel-management/database"
	"hotel-management/models"
	"hotel-management/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Welcome(c *gin.Context) {
	greeting := "Welcome on our hotel management website."
	c.JSON(http.StatusOK, greeting)
}

func getAllEntities[T any](c *gin.Context, db *gorm.DB, entities *[]T) {
	db.Find(entities)
	c.JSON(http.StatusOK, entities)
}

// Afficher toutes les chambres
func GetRooms(c *gin.Context) {
	var chambres []models.Chambre
	getAllEntities(c, database.DB, &chambres)
}

// Afficher les clients
func GetClients(c *gin.Context) {
	var clients []models.Client
	getAllEntities(c, database.DB, &clients)
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

// Nouvelles Reservations
func NewReservation(c *gin.Context) {
	var reservation models.Reservation
	if err := c.ShouldBindJSON(&reservation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Vérifiez la disponibilité de la chambre
	if !services.VerifierDisponibilite(reservation.ChambreID, reservation.DateDebut, reservation.DateFin) {
		c.JSON(http.StatusConflict, gin.H{"error": "Chambre non disponible pour les dates spécifiées"})
		return
	}

	// Récupérez la chambre et calculez le prix total
	var chambre models.Chambre
	if err := database.DB.Where("id = ?", reservation.ChambreID).First(&chambre).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Chambre non trouvée"})
		return
	}

	jours := reservation.DateFin.Sub(reservation.DateDebut).Hours() / 24
	reservation.PrixTotal = chambre.PrixNuit * jours

	// Mettre à jour le statut de la chambre
	chambre.Statut = "occupée"
	database.DB.Save(&chambre)

	// Ajoutez la réservation au client
	var client models.Client
	if err := database.DB.Where("id = ?", reservation.ClientID).First(&client).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Client non trouvé"})
		return
	}
	client.TotalDu += reservation.PrixTotal

	database.DB.Save(&client)
	database.DB.Create(&reservation)

	c.JSON(http.StatusOK, reservation)
}

// Afficher les reservations
func GetReservations(c *gin.Context) {
	var reservations []models.Reservation
	database.DB.Preload("Chambre").Preload("Client").Find(&reservations)
	c.JSON(http.StatusOK, reservations)
}

func UpdateReservation(c *gin.Context) {
	var reservation models.Reservation
	if err := database.DB.Where("id = ?", c.Param("id")).First(&reservation).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Réservation non trouvée"})
		return
	}
	if err := c.ShouldBindJSON(&reservation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Save(&reservation)
	c.JSON(http.StatusOK, reservation)
}

func DeleteReservation(c *gin.Context) {
	var reservation models.Reservation
	if err := database.DB.Where("id = ?", c.Param("id")).First(&reservation).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Réservation non trouvée"})
		return
	}

	var chambre models.Chambre
	if err := database.DB.Where("id = ?", reservation.ChambreID).First(&chambre).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Chambre non trouvée"})
		return
	}

	// Mettre à jour le statut de la chambre
	chambre.Statut = "disponible"
	database.DB.Save(&chambre)

	database.DB.Delete(&reservation)
	c.JSON(http.StatusOK, gin.H{"message": "Réservation supprimée"})
}

func EnregistrerPaiement(c *gin.Context) {
	var input struct {
		ClientID uint    `json:"client_id"`
		Montant  float64 `json:"montant"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var client models.Client
	if err := database.DB.Where("id = ?", input.ClientID).First(&client).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Client non trouvé"})
		return
	}

	client.TotalPaye += input.Montant
	client.TotalDu -= input.Montant

	database.DB.Save(&client)

	c.JSON(http.StatusOK, gin.H{"message": "Paiement enregistré", "solde_restant": client.TotalDu})
}
