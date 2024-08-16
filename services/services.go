package services

import (
	"hotel-management/database"
	"hotel-management/models"
	"time"
)

func VerifierDisponibilite(chambreID uint, dateDebut, dateFin time.Time) bool {
	var reservations []models.Reservation
	database.DB.Where("chambre_id = ? AND date_debut < ? AND date_fin > ?", chambreID, dateFin, dateDebut).Find(&reservations)
	return len(reservations) == 0
}
