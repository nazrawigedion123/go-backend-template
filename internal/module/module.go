package module

import (
	"context"

	"github.com/nazrawigedion123/go-backend-template/internal/constant/db/generated"
)

type SampleModule interface {
	Create(ctx context.Context, params generated.CreateSampleParams) (*generated.Sample, error)

	GetAll(ctx context.Context) ([]generated.Sample, error)

	//

}
