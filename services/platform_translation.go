package services

import (
	"translate-svc/core/logger"
	"translate-svc/models"
	"translate-svc/routes/translations/params"
)

type ServicePlatformTranslation struct {
	Model            models.ModelPlatformTranslation
	ModelTranslation models.ModelTranslation
}

func (s *ServicePlatformTranslation) GetPlatformTranslationByPlatformAndLanguage(platform string, language string) ([]models.EntityPlatformTranslation, error) {
	return s.Model.GetPlatformTranslationByPlatformAndLanguage(platform, language)
}

func (s *ServicePlatformTranslation) AddPlatformTranslations(platformCountry *models.EntityPlatformCountry, translations []params.Translation) (bool, error) {

	for _, trans := range translations {
		pt := models.EntityPlatformTranslation{
			PlatformCountryId: platformCountry.Id,
			Code:              trans.Code,
		}
		err, platformTranslation := s.Model.SavePlatformTranslation(pt)
		if err != nil {
			logger.Log.Infof("SavePlatformTranslation - Error %s", err)
			return false, err
		}

		translation := models.EntityTranslation{
			PlatformTranslationId: platformTranslation.Id,
			Language:              platformCountry.DefaultLanguage,
			Translation:           trans.Translation,
		}
		err, _ = s.ModelTranslation.SaveTranslation(translation)
		if err != nil {
			logger.Log.Infof("SaveTranslation - Error %s", err)
			return false, err
		}
	}

	return true, nil
}
