package services

import (
	"fmt"
	"tys_backend/database"
	"tys_backend/models"
	"tys_backend/services/utils"
	"log"

	"github.com/gofiber/fiber/v2"
)

func GetFeedHunches(c *fiber.Ctx) error {

	//you shouldn't do this by the way, but it's just a demo
	// dbQuery := fmt.Sprintf("SELECT users.user_id, users.user, users.first_name, users.last_name, tweets.tweet, tweets.date_tweet FROM users INNER JOIN tweets ON users.user_id = tweets.user_id INNER JOIN followers ON users.user_id = followers.id_user WHERE followers.id_follower = %s ORDER BY tweets.date_tweet DESC;", c.Params("id"))
	// rows, err := database.DB.Query(dbQuery)

	//avoid the SQL injection by rewriting it like
	dbQuery := "SELECT users.user_id, users.user, users.first_name, users.last_name, hunches.hunch, hunches.date_hunch FROM users INNER JOIN hunches ON users.user_id = hunches.user_id INNER JOIN followers ON users.user_id = followers.id_user WHERE followers.id_follower = ? ORDER BY hunches.date_hunch DESC;"
	rows, err := database.DB.Query(dbQuery, c.Params("id"))

	//check for errors
	if err != nil {
		return utils.DefaultErrorHandler(c, err)
	}
	//close db connection
	defer rows.Close()

	//create a slice of tweets
	var timelineHunches []models.TimelineHunch
	//loop through the result set
	for rows.Next() {
		timelineHunch := models.TimelineHunch{}
		err := rows.Scan(&timelineHunch.User_id, &timelineHunch.User, &timelineHunch.First_name, &timelineHunch.Last_name, &timelineHunch.Hunch, &timelineHunch.Date_hunch)
		if err != nil {
			log.Fatal(err)
		}
		timelineHunches = append(timelineHunches, timelineHunch)
	}
	fmt.Print(timelineHunches)

	utils.ResponseHelperJSON(c, timelineHunches, "timeline", "No timeline found")

	return err
}

func GetFeedHunchesPaginated(c *fiber.Ctx) error {

	// dbQuery := fmt.Sprintf("SELECT users.user_id, users.user, users.first_name, users.last_name, tweets.tweet, tweets.date_tweet FROM users INNER JOIN tweets ON users.user_id = tweets.user_id INNER JOIN followers ON users.user_id = followers.id_user WHERE followers.id_follower = %s ORDER BY tweets.date_tweet DESC LIMIT %s OFFSET %s;", c.Params("id"), c.Params("limit"), c.Params("offset"))
	// avoid a SQL injection by rewriting it like
	dbQuery := "SELECT users.user_id, users.user, users.first_name, users.last_name, hunches.hunch, hunches.date_hunch FROM users INNER JOIN hunches ON users.user_id = hunches.user_id INNER JOIN followers ON users.user_id = followers.id_user WHERE followers.id_follower = ? ORDER BY hunches.date_hunch DESC LIMIT ? OFFSET ?;"

	rows, err := database.DB.Query(dbQuery, c.Params("id"), c.Params("limit"), c.Params("offset"))
	if err != nil {
		return utils.DefaultErrorHandler(c, err)
	}

	defer rows.Close()

	var timelineHunches []models.TimelineHunch
	for rows.Next() {
		timelineHunch := models.TimelineHunch{}
		err := rows.Scan(&timelineHunch.User_id, &timelineHunch.User, &timelineHunch.First_name, &timelineHunch.Last_name, &timelineHunch.Hunch, &timelineHunch.Date_hunch)
		if err != nil {
			log.Fatal(err)
		}
		timelineHunches = append(timelineHunches, timelineHunch)
	}
	//TODO: implement a response with pages and all that pagination jazz
	utils.ResponseHelperJSON(c, timelineHunches, "timeline", "No timeline found")

	return err
}
