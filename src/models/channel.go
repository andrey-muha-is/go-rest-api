package models

type Channel struct {
	ID    string `json:"id" db:"id"`
	Title string `json:"title" db:"title"`
	Icon  string `json:"icon" db:"icon"`
}

type ChannelsResponse struct {
	Channels []Channel `json:"channels"`
}