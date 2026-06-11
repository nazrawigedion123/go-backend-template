package initiator

import (
	"github.com/OnePulseOmni/pulse-wallet/internal/constant/model/persistencedb"
	"github.com/OnePulseOmni/pulse-wallet/internal/storage"
	samplestorage "github.com/OnePulseOmni/pulse-wallet/internal/storage/sample-storage"
	"github.com/OnePulseOmni/pulse-wallet/platform/logger"
)

type Persistance struct {
	Sample storage.Sample
	Logger logger.Logger
}

func initPersistence(persistencedb *persistencedb.PersistenceDB, logger logger.Logger) *Persistance {
	sample := samplestorage.New(logger, persistencedb)

	return &Persistance{
		Sample: sample,
		Logger: logger,
	}

}
