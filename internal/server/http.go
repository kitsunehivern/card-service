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
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	})

	cardHdl := handler.NewCardHandler(cardSvc)
	r.POST("/card/request", cardHdl.RequestCard)
	r.GET("/card/get/:id", cardHdl.GetCard)
	r.POST("/card/activate", cardHdl.ActivateCard)
	r.POST("/card/block", cardHdl.BlockCard)
	r.POST("/card/unblock", cardHdl.UnblockCard)
	r.POST("/card/retire", cardHdl.RetireCard)
	r.POST("/card/close", cardHdl.CloseCard)

	return r
}
