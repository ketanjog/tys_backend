package api

import (
	"tys_backend/services"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	api := app.Group("/api")

	//get all unordened users
	api.Get("/users", services.GetUsers)
	//get all users ordered by age ASC
	api.Get("/users/age", services.GetUsersByAgeAsc)

	//get all unordened hunches from db
	api.Get("/hunches", services.GetHunches)

	//http://localhost:3000/api/feed/1
	//get MOST RECENT feed/timeline for the user
	api.Get("/feed/:id", services.GetFeedHunches)
	//pagination
	api.Get("/feed/:id/:limit/:offset", services.GetFeedHunchesPaginated)
	//can try https://github.com/gofiber/fiber/issues/193#issuecomment-591976894
	//a whole presentation on why you shouldn't do what I did:
	//https://use-the-index-luke.com/no-offset

}
