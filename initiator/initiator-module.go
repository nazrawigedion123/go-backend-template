package initiator

import (
	"github.com/nazrawigedion123/go-backend-template/internal/module"
	samplemodule "github.com/nazrawigedion123/go-backend-template/internal/module/sample-Module"
	"github.com/nazrawigedion123/go-backend-template/platform/logger"
)

type Module struct {
	Sample module.SampleModule
	Logger logger.Logger
}

func initModule(persistance *Persistance, logger logger.Logger) *Module {
	sample := samplemodule.New(logger, persistance.Sample)
	return &Module{
		Sample: sample,
		Logger: logger,
	}

}
