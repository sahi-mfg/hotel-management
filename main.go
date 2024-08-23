package main

import (
	"hotel-management/controllers"
	"hotel-management/database"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	database.InitDatabase()

	r.GET("/", controllers.Welcome)

	// Routes pour les chambres
	r.GET("/chambres", controllers.GetRooms)

	// Routes pour les reservations
	r.GET("/reservations", controllers.GetReservations)
	r.POST("/reservations", controllers.NewReservation)
	r.DELETE("/reservations", controllers.DeleteReservation)

	// Routes pour les clients
	r.GET("/clients", controllers.GetClients)
	r.POST("/clients", controllers.NewClient)
	r.DELETE("/clients", controllers.DeleteClient)

	r.Run() // Démarre le serveur sur localhost:8080
}
