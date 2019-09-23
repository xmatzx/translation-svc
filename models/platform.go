package models

import (
	"translate-svc/core/db"
)

type Platform struct {
	Id   int    `gorm:"column:id"`
	Name string `gorm:"column:name"`
	Code string `gorm:"column:code"`
}

func (m *Platform) GetPlatformByCode(code string) (*Platform, error) {
	var platform Platform

	dbConn := db.DB
	err := dbConn.Where(Platform{Code: code}).First(&platform).Error

	if err != nil {
		return nil, err
	}

	return &platform, err
}
