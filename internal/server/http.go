package server

import (
	"card-service/internal/handler"
	"card-service/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(cardSvc *service.CardService) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	cardHdl := handler.NewCardHandler(cardSvc)
	r.POST("/card/request", cardHdl.RequestCard)
	r.GET("/card/:id", cardHdl.GetCard)
	r.POST("/card/:id/activate", cardHdl.ActivateCard)
	r.POST("/card/:id/block", cardHdl.BlockCard)
	r.POST("/card/:id/unblock", cardHdl.UnblockCard)
	r.POST("/card/:id/retire", cardHdl.RetireCard)
	r.POST("/card/:id/close", cardHdl.CloseCard)

	return r
}
