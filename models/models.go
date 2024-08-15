package models

type Client struct {
	ID        uint   `gorm:"primaryKey"`
	Nom       string `gorm:"size:100;not null"`
	Prenom    string `gorm:"size:100;not null"`
	Telephone string `gorm:"size:20"`
	Email     string `gorm:"size:100"`
}

type Chambre struct {
	ID          uint   `gorm:"primaryKey"`
	Numero      string `gorm:"size:10;not null;unique"`
	TypeChambre string `gorm:"size:50"`
	PrixNuit    float64
	Statut      string `gorm:"size:50;default:'disponible'"`
}

type Reservation struct {
	ID          uint `gorm:"primaryKey"`
	ClientID    uint `gorm:"not null"`
	ChambreID   uint `gorm:"not null"`
	DateArrivee string
	DateDepart  string
	Statut      string `gorm:"size:50;default:'confirmée'"`
}

type Occupation struct {
	ID            uint `gorm:"primaryKey"`
	ReservationID uint `gorm:"not null"`
	DateCheckin   string
	DateCheckout  string
	Statut        string `gorm:"size:50;default:'en cours'"`
}
