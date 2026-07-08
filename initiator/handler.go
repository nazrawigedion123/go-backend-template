package initiator

import (
	"github.com/nazrawigedion123/go-backend-template/internal/handler"
	"github.com/nazrawigedion123/go-backend-template/internal/handler/samplehandler"
	"github.com/nazrawigedion123/go-backend-template/platform/logger"
)

type Handler struct {
	SampleHandler handler.SampleHandler
}

func initHandler(module *Module, logger logger.Logger) *Handler {
	sample := samplehandler.New(logger, module.Sample)
	return &Handler{
		SampleHandler: sample,
	}
}
