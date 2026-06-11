package module

import (
	"context"

	"github.com/OnePulseOmni/pulse-wallet/internal/constant/model/db"
)

type SampleModule interface {
	Create(ctx context.Context, params db.CreateSampleParams) (*db.Sample, error)

	GetAll(ctx context.Context) ([]db.Sample, error)

	//

}
