package repositories

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	m "../models"
)

type ChannelsRepository struct {
	db *sqlx.DB
	tableName string
}

func NewChannelsRepository(db *sqlx.DB, tableName string) *ChannelsRepository {
	return &ChannelsRepository{db, tableName}
}

func (c *ChannelsRepository) FindAll() (*m.ChannelsResponse, error) {
	channels := []m.Channel{}
	
	query := fmt.Sprintf(`SELECT id,title,icon FROM %s`, c.tableName)
	err := c.db.Select(&channels, query)
	
	if err != nil {
		return nil, err
	}

	channelsResponse := m.ChannelsResponse{channels} 

	return &channelsResponse, nil
}

func (c *ChannelsRepository) FindById(id string) (*m.Channel, error) {
	channel := m.Channel{}

	query := fmt.Sprintf(`SELECT id,title,icon FROM %s WHERE id=?`, c.tableName)
	err := c.db.Get(&channel, query, id)

	if err != nil {
		return nil, err
	}

	return &channel, nil
}
