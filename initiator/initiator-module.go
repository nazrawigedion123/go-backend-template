package initiator

import (
	"github.com/OnePulseOmni/pulse-wallet/internal/module"
	samplemodule "github.com/OnePulseOmni/pulse-wallet/internal/module/sample-Module"
	"github.com/OnePulseOmni/pulse-wallet/platform/logger"
)



type Module struct{
	Sample module.SampleModule
	Logger logger.Logger
}

func initModule(persistance *Persistance, logger logger.Logger) *Module{
	sample:=samplemodule.New(logger,persistance.Sample)
	return &Module{
		Sample: sample,
		Logger: logger,
	}

}