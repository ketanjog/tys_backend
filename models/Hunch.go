package models

type Hunch struct {
	Id        int    `json:"id"`
	User_id   int    `json:"user_id"`
	Hunch     string `json:"hunch"`
	Date_hunch string `json:"date_hunch"`
}

type Hunches struct {
	Hunches []Hunch `json:"hunches"`
}