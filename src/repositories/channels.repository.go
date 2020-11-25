package repositories

import (
	"fmt"

	m "../models"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type ChannelsRepository struct {
	db        *sqlx.DB
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
		zap.S().Errorf("Error fetching channels. Error: %s", err)
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
		zap.S().Errorf("Error fetching channel with id = %s. Error: %s", id, err)
		return nil, err
	}

	return &channel, nil
}
