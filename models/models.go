package models

import (
	"time"

	"gorm.io/gorm"
)

type Client struct {
	gorm.Model
	Nom          string        `json:"nom"`
	Prenom       string        `json:"prenom"`
	Telephone    string        `json:"telephone"`
	Reservations []Reservation `json:"reservations"` // Relation avec les réservations
}

type Chambre struct {
	gorm.Model
	Numero      string  `json:"numero"`
	TypeChambre string  `json:"type_chambre"`
	PrixNuit    float64 `json:"prix_nuit"`
	Statut      string  `json:"statut"` // disponible, occupée, etc.
}

type Reservation struct {
	gorm.Model
	DateDebut time.Time `json:"date_debut"`
	DateFin   time.Time `json:"date_fin"`
	ChambreID uint      `json:"chambre_id"`
	ClientID  uint      `json:"client_id"`
	PrixTotal float64   `json:"prix_total"`
}
