package samplerouting

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nazrawigedion123/go-backend-template/internal/glue/routing"
	"github.com/nazrawigedion123/go-backend-template/internal/handler"
	"github.com/nazrawigedion123/go-backend-template/platform/logger"
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
