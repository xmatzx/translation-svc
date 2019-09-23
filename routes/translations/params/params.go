package params

// Binding from JSON
type TranslationRequest struct {
	Country  string `form:"country" json:"country" binding:"required,len=2"`
	Platform string `form:"platform" json:"platform" binding:"required"`
	Language string `form:"lang" json:"lang" binding:"required,len=2"`
}

type TranslationImportRequest struct {
	Country      string        `form:"country" json:"country" binding:"required,len=2"`
	Platform     string        `form:"platform" json:"platform" binding:"required"`
	Translations []Translation `form:"translations" json:"translations" binding:"required"`
}

type Translation struct {
	Code        string `form:"code" json:"code" binding:"required"`
	Translation string `form:"translation" json:"translation" binding:"required"`
}
