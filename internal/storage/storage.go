package storage

import (
	"context"

	"github.com/nazrawigedion123/go-backend-template/internal/constant/model/db"
)

type Sample interface {
	Create(ctx context.Context, params db.CreateSampleParams) (*db.Sample, error)

	GetAll(ctx context.Context) ([]db.Sample, error)
}
