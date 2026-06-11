package initiator

import (
	"github.com/OnePulseOmni/pulse-wallet/internal/handler"
	samplehandler "github.com/OnePulseOmni/pulse-wallet/internal/handler/sample-Handler"
	"github.com/OnePulseOmni/pulse-wallet/platform/logger"
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
