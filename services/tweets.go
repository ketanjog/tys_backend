package services

import (
	"tys_backend/database"
	"tys_backend/models"
	"log"

	// "database/sql"
	"fmt"

	"tys_backend/services/utils"

	"github.com/gofiber/fiber/v2"
)

func GetHunches(c *fiber.Ctx) error {

	//get the tweets from db
	res, err := database.DB.Query("SELECT * FROM hunches")
	//check for errors
	// if err != nil {
	// 	c.Status(500).JSON(&fiber.Map{
	// 		"success": false,
	// 		"error":   err,
	// 	})
	// 	return nil
	// }
	if err != nil {
		return utils.DefaultErrorHandler(c, err)
	}

	//close the result set
	defer res.Close()
	//create a slice of tweets
	var hunches []models.Hunch
	//loop through the result set
	for res.Next() {

		hunch := models.Hunch{}

		err := res.Scan(&hunch.Id, &hunch.User_id, &hunch.Hunch, &hunch.Date_hunch)

		if err != nil {
			log.Fatal(err)
		}
		//best logger amirite
		fmt.Printf("%v\n", hunch)
		//storing the data in a slice
		hunches = append(hunches, hunch)
	}

	// fmt.Printf("%v\n", tweets)

	//send the tweets to the client
	// if tweets != nil {
	// 	c.Status(200).JSON(&fiber.Map{
	// 		"success": true,
	// 		"tweets":  tweets,
	// 	})
	// } else {
	// 	c.Status(404).JSON(&fiber.Map{
	// 		"success": false,
	// 		"error":   "No tweets found",
	// 	})
	// }
	utils.ResponseHelperJSON(c, hunches, "hunches", "No hunches found")

	return err
}
