package initiator

import (
	"github.com/nazrawigedion123/go-backend-template/internal/constant/db/dbinterface"
	"github.com/nazrawigedion123/go-backend-template/internal/storage"
	"github.com/nazrawigedion123/go-backend-template/internal/storage/samplestorage"
	"github.com/nazrawigedion123/go-backend-template/platform/logger"
)

type Persistance struct {
	Sample storage.Sample
	Logger logger.Logger
}

func initPersistence(persistencedb *dbinterface.PersistenceDB, logger logger.Logger) *Persistance {
	sample := samplestorage.New(logger, persistencedb)

	return &Persistance{
		Sample: sample,
		Logger: logger,
	}

}
