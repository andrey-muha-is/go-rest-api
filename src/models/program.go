package models

type Program struct {
	Started     int    `json:"started" db:"started"`
	Ended       int    `json:"ended" db:"ended"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
}

type ProgramsResponse struct {
	Programs []Program `json:"programs"`
}