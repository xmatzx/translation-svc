package services

import (
	"translate-svc/models"
)

type ServiceTranslation struct {
	Model models.ModelTranslation
}

func (s *ServiceTranslation) GetTranslationByPlatformLanguage(platform string, language string) ([]models.EntityTranslation, error) {
	return s.Model.GetTranslationByPlatformLanguage(platform, language)
}

func (s *ServiceTranslation) SaveTranslation(translation models.EntityTranslation) (error, *models.EntityTranslation) {
	return s.Model.SaveTranslation(translation)
}
