package services

import "translate-svc/models"

type ServicePlatformCountry struct {
	Model models.ModelPlatformCountry
}

func (s *ServicePlatformCountry) GetPlatformCountry(platform string, country string) (*models.EntityPlatformCountry, error) {
	return s.Model.GetPlatformCountry(platform, country)
}
