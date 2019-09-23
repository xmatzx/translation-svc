package models

import (
	"translate-svc/core/db"
	"translate-svc/core/logger"
)

const (
	tablePlatformTranslation = "platform_translations"
)

type EntityPlatformTranslation struct {
	Id                int    `gorm:"column:id"`
	PlatformCountryId int    `gorm:"column:platform_country_id"`
	Code              string `gorm:"column:code"`
	Language          string `gorm:"-"`
	Translation       string `gorm:"-"`
}

type ModelPlatformTranslationInterface interface {
	GetPlatformTranslationByPlatformAndLanguage(platform string, language string) ([]EntityPlatformTranslation, error)
	SavePlatformTranslation(platformTranslation EntityPlatformTranslation) (error, *EntityPlatformTranslation)
}

type ModelPlatformTranslation struct {
}

func (ModelPlatformTranslation) TableName() string {
	return tablePlatformTranslation
}

func (m *ModelPlatformTranslation) GetPlatformTranslationByPlatformAndLanguage(platform string, language string) ([]EntityPlatformTranslation, error) {
	var translations []EntityPlatformTranslation

	dbConn := db.DB
	err := dbConn.
		Table("platform_translations").
		Select("" +
			"platform_translations.id, " +
			"platform_translations.platform_country_id, " +
			"platform_translations.code, " +
			"translations.language, " +
			"translations.translation").
		Joins("JOIN platforms_countries ON platforms_countries.id = platform_translations.platform_country_id").
		Joins("JOIN platforms ON platforms.id = platforms_countries.platform_id").
		Joins("LEFT JOIN translations ON (translations.platform_translation_id = platform_translations.id AND translations.language = ?)", language).
		Where("platforms.code = ? ", platform).
		Scan(&translations).
		Error

	if err != nil {
		return nil, err
	}

	return translations, nil
}

func (m *ModelPlatformTranslation) SavePlatformTranslation(platformTranslation EntityPlatformTranslation) (error, *EntityPlatformTranslation) {
	err := db.DB.
		Table("platform_translations").
		Where("platform_translations.platform_country_id = ? ", platformTranslation.PlatformCountryId).
		Where("platform_translations.code = ? ", platformTranslation.Code).
		//Where(EntityPlatformTranslation{PlatformCountryId: platformTranslation.PlatformCountryId, Code: platformTranslation.Code}).
		FirstOrCreate(&platformTranslation).
		Error

	if err != nil {
		return err, nil
	}

	logger.Log.Infof("SavePlatformTranslation - PlatformTranslationId %d", platformTranslation.Id)
	return nil, &platformTranslation
}
