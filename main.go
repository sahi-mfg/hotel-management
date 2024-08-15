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
	r.GET("/chambres", controllers.GetChambres)
	r.POST("/chambres", controllers.CreateChambre)

	r.GET("/reservations", controllers.GetRservations)
	r.POST("/reservations", controllers.NewReservation)
	r.DELETE("/reservations", controllers.DeleteReservation)

	r.GET("/clients", controllers.GetClients)
	r.POST("/clients", controllers.NewClient)
	r.DELETE("/clients", controllers.DeleteClient)

	r.Run() // Démarre le serveur sur localhost:8080
}
