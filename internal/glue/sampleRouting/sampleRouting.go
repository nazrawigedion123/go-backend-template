package samplerouting

import (
	"net/http"

	"github.com/OnePulseOmni/pulse-wallet/internal/glue/routing"
	"github.com/OnePulseOmni/pulse-wallet/internal/handler"
	"github.com/OnePulseOmni/pulse-wallet/platform/logger"
	"github.com/gin-gonic/gin"
)

func RegisterSampleRouter(
	group *gin.RouterGroup,
	sampleHandler handler.SampleHandler,
	logger logger.Logger,

) {

	payments := []routing.Route{
		{
			Method:     http.MethodPost,
			Path:       "/api/v1/sample-handlers",
			Handler:    sampleHandler.Create,
			Middleware: []gin.HandlerFunc{},
		}, {
			Method:     http.MethodGet,
			Path:       "/api/v1/sample-handlers",
			Handler:    sampleHandler.GetAll,
			Middleware: []gin.HandlerFunc{},
		},
	}

	routing.RegisterRoute(group, payments, logger)
}
