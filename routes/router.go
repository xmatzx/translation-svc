package routes

import (
	"github.com/gin-gonic/gin"
	"translate-svc/routes/translations"
)

func RegisterRoutes(r *gin.Engine) {
	translationController := translations.New()

	r.GET("translations/all", translationController.Handle)

	r.POST("translations/import", translationController.Import)

	// However, this one will match /user/john/ and also /user/john/send
	// If no other routers match /user/john, it will redirect to /user/john/
	//router.GET("/translation/:name/*action", func(c *gin.Context) {
	//	name := c.Param("name")
	//	action := c.Param("action")
	//	message := name + " is " + action
	//	c.String(http.StatusOK, message)
	//})

	//r.GET("/translation", func(c *gin.Context) {
	//	country := c.DefaultQuery("country", "Guest")
	//	platform := c.Query("platform") // shortcut for c.Request.URL.Query().Get("lastname")
	//	lang := c.Query("lang") // shortcut for c.Request.URL.Query().Get("lastname")
	//
	//	c.JSON(200, gin.H{
	//		"country":  country,
	//		"platform": platform,
	//		"language":    lang,
	//	})
	//})
}
