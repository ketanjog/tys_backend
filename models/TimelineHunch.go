package models

type TimelineHunch struct {
	User_id int `json:"user_id"`
	User string `json:"user"`
	First_name string `json:"first_name"`
	Last_name string `json:"last_name"`
	Hunch string `json:"hunch"`
	Date_hunch string `json:"date_hunch"`
}

type TimelineHunches struct {
	Timeline_hunches []TimelineHunch `json:"timeline_hunches"`
}