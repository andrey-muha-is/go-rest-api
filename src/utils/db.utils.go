package utils

import (
	"fmt"

	"../config"
)

func GetDataSourceName(c *config.AppConfig) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		c.DbUser,
		c.DbPassword,
		c.DbHost,
		c.DbPort,
		c.DbName)
}
