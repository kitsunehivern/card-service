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
	r.PATCH("/card/:id/activate", cardHdl.ActivateCard)
	r.PATCH("/card/:id/block", cardHdl.BlockCard)
	r.PATCH("/card/:id/unblock", cardHdl.UnblockCard)
	r.PATCH("/card/:id/close", cardHdl.CloseCard)

	return r
}
