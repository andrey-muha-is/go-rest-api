package models

import (
	"github.com/guregu/null"
)

type Program struct {
	Started     int         `json:"started" db:"started"`
	Ended       int         `json:"ended" db:"ended"`
	Title       string      `json:"title" db:"title"`
	Description null.String `json:"description" db:"description"`
}

type ProgramsResponse struct {
	Programs []Program `json:"programs"`
}
