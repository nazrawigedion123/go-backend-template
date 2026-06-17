package samplestorage

import (
	"context"
	"fmt"

	dbinterface "github.com/nazrawigedion123/go-backend-template/internal/constant/db/db_interface"
	"github.com/nazrawigedion123/go-backend-template/internal/constant/db/generated"
	"github.com/nazrawigedion123/go-backend-template/internal/storage"
	"github.com/nazrawigedion123/go-backend-template/platform/logger"
	"go.uber.org/zap"
)

type sampleStorage struct {
	logger        logger.Logger
	persistencedb *dbinterface.PersistenceDB
}

// New creates a new sample storage instance
func New(logger logger.Logger, persistencedb *dbinterface.PersistenceDB) storage.Sample {
	return &sampleStorage{
		logger:        logger,
		persistencedb: persistencedb,
	}
}

// Create inserts a new sample record
func (s *sampleStorage) Create(ctx context.Context, params generated.CreateSampleParams) (*generated.Sample, error) {

	sample, err := s.persistencedb.Queries.CreateSample(ctx, params)
	if err != nil {

		return nil, fmt.Errorf("create sample: %w", err)
	}

	return &sample, nil
}

// GetAll retrieves all samples
func (s *sampleStorage) GetAll(ctx context.Context) ([]generated.Sample, error) {

	samples, err := s.persistencedb.Queries.GetAllSamples(ctx)
	if err != nil {
		s.logger.Error(ctx, "failed to get all samples", zap.Error(err))
		return nil, fmt.Errorf("get all samples: %w", err)
	}
	return samples, nil
}

//
