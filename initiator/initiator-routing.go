package initiator

import (
	"github.com/gin-gonic/gin"
	samplerouting "github.com/nazrawigedion123/go-backend-template/internal/glue/sampleRouting"
	"github.com/nazrawigedion123/go-backend-template/platform/logger"
)

func initRoute(grg *gin.RouterGroup, handler *Handler, logger logger.Logger) {
	samplerouting.RegisterSampleRouter(grg, handler.SampleHandler, logger)
}
