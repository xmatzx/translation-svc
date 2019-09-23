package models

import (
	"github.com/jinzhu/gorm"
	"translate-svc/core/db"
	"translate-svc/core/logger"
)

const (
	tablePlatformCountry = "platforms_countries"
)

type EntityPlatformCountry struct {
	Id              int    `gorm:"column:id"`
	PlatformId      int    `gorm:"column:platform_id"`
	CountryId       int    `gorm:"column:country_id"`
	DefaultLanguage string `gorm:"column:default_language"`
}

type ModelPlatformCountryInterface interface {
	GetPlatformCountry(platform string, country string) (*EntityPlatformCountry, error)
}

type ModelPlatformCountry struct {
}

func (ModelPlatformCountry) TableName() string {
	return tablePlatformCountry
}

func (m *ModelPlatformCountry) GetPlatformCountry(platform string, country string) (*EntityPlatformCountry, error) {
	var platformCountry EntityPlatformCountry

	dbConn := db.DB
	err := dbConn.
		Table("platforms_countries").
		Joins("JOIN countries ON countries.id = platforms_countries.country_id").
		Joins("JOIN platforms ON platforms.id = platforms_countries.platform_id").
		Where("countries.iso_code = ?", country).
		Where("platforms.code = ?", platform).
		First(&platformCountry).
		Error

	if err != nil {
		return nil, err
	}

	if gorm.IsRecordNotFoundError(err) {
		logger.Log.Infof("GetPlatformCountry - platform: %s country: %s - %s", platform, country, err)
		return nil, nil
	}

	return &platformCountry, nil
}
