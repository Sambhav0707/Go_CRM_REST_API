package lead

import (
	"github.com/Sambhav0707/Go_CRM/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Lead struct {
	gorm.Model
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
}

func GetLeads(c *fiber.Ctx) {
	//create a empty list of Lead
	var leads []Lead

	//find the leads in the database
	db := database.DBConn
	db.Find(&leads)

	//send the req
	c.JSON(leads)

}

func GetLead(c *fiber.Ctx) {
	//get the id from the params
	id := c.Params("id")

	//create a lead
	var lead Lead

	//search in the database
	db := database.DBConn

	db.Find(&lead, id)

	// send the response

	c.JSON(lead)

}

func CreateLead(c *fiber.Ctx) {
	// get the db
	db := database.DBConn

	lead := new(Lead)

	if err := c.BodyParser(lead); err != nil {
		c.Status(503).Send(err)
		return
	}

	db.Create(&lead)
	c.JSON(lead)
}

func DeleteLead(c *fiber.Ctx) {
	db := database.DBConn

	id := c.Params("id")

	var lead Lead

	db.First(&lead, id)

	if lead.Name == "" {
		c.Status(503).Send("no lead found in the db")
		return
	}

	db.Delete(&lead)

	c.Send("Lead deleted successfully")
}
