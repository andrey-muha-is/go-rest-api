package repositories

import (
	"fmt"
	"github.com/jmoiron/sqlx"

	m "../models"
)

type ProgramsRepository struct {
	db        *sqlx.DB
	tableName string
}

func NewProgramsRepository(db *sqlx.DB, tableName string) *ProgramsRepository {
	return &ProgramsRepository{db, tableName}
}

func (p *ProgramsRepository) FindAll() (*m.ProgramsResponse, error) {
	programs := []m.Program{}
	
	query := fmt.Sprintf(`SELECT started,ended,title,description FROM %s`, p.tableName)
	err := p.db.Select(&programs, query)
	
	if err != nil {
		return nil, err
	}

	programsResponse := m.ProgramsResponse{programs} 

	return &programsResponse, nil
}

func (p *ProgramsRepository) FindByChannelId(channelID string) (*m.ProgramsResponse, error) {
	programs := []m.Program{}

	query := fmt.Sprintf(`SELECT started,ended,title,description FROM %s WHERE channel_id=?`, p.tableName)
	err := p.db.Select(&programs, query, channelID)

	if err != nil {
		return nil, err
	}

	programsResponse := m.ProgramsResponse{programs}

	return &programsResponse, nil
}
