package translations

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"translate-svc/routes/translations/params"

	"translate-svc/core/logger"
	"translate-svc/models"
	"translate-svc/services"
)

type TranslationController struct{}

func New() *TranslationController {
	return &TranslationController{}
}

type PlatformTranslationsResponse struct {
	Id          int    `json:"id"`
	Code        string `json:"code"`
	Translation string `json:"translation"`
}

func (h *TranslationController) Handle(c *gin.Context) {

	var request params.TranslationRequest
	if err := c.BindQuery(&request); err != nil {
		logger.Log.Info("Invalid request params:", request)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var modelPlatformTranslation models.ModelPlatformTranslation
	translationService := &services.ServicePlatformTranslation{Model: modelPlatformTranslation}

	translations, err := translationService.GetPlatformTranslationByPlatformAndLanguage(
		request.Platform,
		request.Language,
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"translations": buildPlatformTranslationsResponse(translations),
	})
}

func (h *TranslationController) Import(c *gin.Context) {

	var request params.TranslationImportRequest
	if err := c.BindJSON(&request); err != nil {
		logger.Log.Info("Invalid request params:", request)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var modelPlatformCountry models.ModelPlatformCountry
	platformCountryService := &services.ServicePlatformCountry{Model: modelPlatformCountry}

	platformCountry, err := platformCountryService.GetPlatformCountry(
		request.Platform,
		request.Country,
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if platformCountry == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "PlatformCountry not found"})
		return
	}

	var modelPlatformTranslation models.ModelPlatformTranslation
	var modelTranslation models.ModelTranslation
	platformTranslationService := &services.ServicePlatformTranslation{
		Model:            modelPlatformTranslation,
		ModelTranslation: modelTranslation,
	}

	result, err := platformTranslationService.AddPlatformTranslations(platformCountry, request.Translations)

	c.JSON(200, gin.H{
		"status": result,
	})
}

func buildPlatformTranslationsResponse(translations []models.EntityPlatformTranslation) []PlatformTranslationsResponse {
	var result []PlatformTranslationsResponse
	for _, t := range translations {
		result = append(result, PlatformTranslationsResponse{t.Id, t.Code, t.Translation})
	}

	return result
}
