package initiator

import (
	samplerouting "github.com/OnePulseOmni/pulse-wallet/internal/glue/sampleRouting"
	"github.com/OnePulseOmni/pulse-wallet/platform/logger"
	"github.com/gin-gonic/gin"
)

func initRoute(grg *gin.RouterGroup, handler *Handler, logger logger.Logger) {
	samplerouting.RegisterSampleRouter(grg, handler.SampleHandler, logger)
}
