package services

import "translate-svc/models"

type ServicePlatform struct {
	Model models.Platform
}

func (s *ServicePlatform) GetPlatformByCode(code string) (*models.Platform, error) {
	return s.Model.GetPlatformByCode(code)
}
