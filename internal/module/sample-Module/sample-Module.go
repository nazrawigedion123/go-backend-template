package samplemodule

import (
	"context"
	"fmt"

	"github.com/nazrawigedion123/go-backend-template/internal/constant/db/generated"
	"github.com/nazrawigedion123/go-backend-template/internal/module"
	"github.com/nazrawigedion123/go-backend-template/internal/storage"
	"github.com/nazrawigedion123/go-backend-template/platform/logger"
	"go.uber.org/zap"
)

type sampleModule struct {
	sampleStorage storage.Sample
	logger        logger.Logger
}

func New(logger logger.Logger, sampleStorage storage.Sample) module.SampleModule {
	return &sampleModule{
		logger:        logger,
		sampleStorage: sampleStorage,
	}
}

func (m *sampleModule) Create(ctx context.Context, params generated.CreateSampleParams) (*generated.Sample, error) {
	m.logger.Debug(ctx, "creating new sample",
		zap.String("name", params.Name),
		zap.String("email", params.Email),
	)
	sample, err := m.sampleStorage.Create(ctx, params)
	if err != nil {
		m.logger.Error(ctx, "failed to create sample",
			zap.Error(err),
			// zap.String("params", params),
		)
		return nil, err
	}

	m.logger.Info(ctx, "sample created successfully",
		zap.String("sample_id", string(sample.ID)),
		zap.String("email", sample.Email),
	)

	return sample, nil

}

func (m *sampleModule) GetAll(ctx context.Context) ([]generated.Sample, error) {
	m.logger.Debug(ctx, "fetching all samples")

	samples, err := m.sampleStorage.GetAll(ctx)
	if err != nil {
		m.logger.Error(ctx, "failed to get all samples", zap.Error(err))
		return nil, fmt.Errorf("get all samples: %w", err)
	}

	m.logger.Debug(ctx, "samples fetched successfully", zap.String("count", string(len(samples))))
	return samples, nil
}

//
