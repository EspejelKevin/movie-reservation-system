package database

import "gorm.io/gorm"

type Connection struct {
	db *gorm.DB
}

func NewConnection(dialector gorm.Dialector) (*Connection, error) {
	db, err := gorm.Open(dialector, &gorm.Config{})

	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()

	if err != nil {
		return nil, err
	}

	err = sqlDB.Ping()

	if err != nil {
		return nil, err
	}

	return &Connection{db}, nil
}

func (connection *Connection) GetDB() *gorm.DB {
	return connection.db
}
