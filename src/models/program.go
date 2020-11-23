package models

type Program struct {
	Started     int    `json:"started"`
	Ended       int    `json:"ended"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
