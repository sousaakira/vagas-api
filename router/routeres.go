package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sousaakira/vagas-api.git/handler"
)

func initializeRouters(router *gin.Engine) {

	v1 := router.Group("/api/v1")
	{
		// Show Opening
		v1.GET("/opening", handler.ListOpeningHandler)

		v1.POST("/opening", handler.CreateOpeningHandler)

		v1.DELETE("/opening", handler.DeleteOpeningHandler)

		v1.PUT("/opening", handler.UpdateOpeningHandler)

		v1.GET("/openings", handler.ListOpeningsHandler)
	}

}
