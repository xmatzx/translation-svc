package models

import (
	"translate-svc/core/db"
	"translate-svc/core/logger"
)

const (
	tableTranslation = "translations"
)

type ModelTranslationInterface interface {
	GetTranslationByPlatformLanguage(platform string, language string) (*EntityTranslation, error)
	SaveTranslation(translation EntityTranslation) (error, *EntityTranslation)
}

type EntityTranslation struct {
	Id                    int    `gorm:"column:id"`
	PlatformTranslationId int    `gorm:"column:platform_translation_id"`
	Code                  string `gorm:"-"`
	Language              string `gorm:"column:language"`
	Translation           string `gorm:"column:translation"`
}

type ModelTranslation struct {
}

func (EntityTranslation) TableName() string {
	return tableTranslation
}

func (m *ModelTranslation) GetTranslationByPlatformLanguage(platform string, language string) ([]EntityTranslation, error) {
	var translations []EntityTranslation

	dbConn := db.DB
	err := dbConn.
		Table("translations").
		Select("").
		Joins("JOIN platform_translations ON platform_translations.id = translations.platform_translation_id").
		Joins("JOIN platforms_countries ON platforms_countries.id = platform_translations.platform_country_id").
		Joins("JOIN platforms ON platforms.id = platforms_countries.platform_id").
		Where("platforms.code = ? ", platform).
		Where("translations.language = ?", language).
		Scan(&translations).
		Error

	if err != nil {
		return nil, err
	}

	return translations, nil
}

func (m *ModelTranslation) SaveTranslation(translation EntityTranslation) (error, *EntityTranslation) {
	err := db.DB.Model(&EntityTranslation{}).Create(&translation).Error
	if err != nil {
		return err, nil
	}

	logger.Log.Infof("SaveTranslation - TranslationId %d", translation.Id)
	return nil, &translation
}
