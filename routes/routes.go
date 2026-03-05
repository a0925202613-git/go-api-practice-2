package routes

import (
	"go-api-practice-2/handlers"

	"github.com/gin-gonic/gin"
)

// Setup 註冊 API 路由（只保留一組給練習）
//
// 主題：三麗鷗「人魚漢頓」的夥伴（漢頓的朋友，如小百合等）
// Gin：r.Group 為路徑前綴；:id 為路徑參數，在 handler 用 c.Param("id") 取得
func Setup(r *gin.Engine) {
	api := r.Group("/api")
	{
		// 人魚漢頓的夥伴（GET / POST / PUT / DELETE 練習用）
		api.GET("/partners", handlers.GetPartners)
		api.GET("/partners/:id", handlers.GetPartnerByID)
		api.POST("/partners", handlers.CreatePartner)
		api.PUT("/partners/:id", handlers.UpdatePartner)
		api.DELETE("/partners/:id", handlers.DeletePartner)
	}
}
