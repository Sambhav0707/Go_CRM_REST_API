package main

import (
	"fmt"

	"github.com/Sambhav0707/Go_CRM/database"
	"github.com/Sambhav0707/Go_CRM/lead"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// function to set up routed
func setUpRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.CreateLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}

// funciotn to connect to the database
func databaseInit() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "lead.db")

	if err != nil {
		panic("there was error connecting to the database")
	}
	fmt.Println("Connnection to the databse established successfully")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("database migrated successfully")
}

func main() {
	app := fiber.New()
	databaseInit()
	setUpRoutes(app)
	app.Listen(3000)
	defer database.DBConn.Close()

}
