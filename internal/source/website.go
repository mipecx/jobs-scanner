package source

import (
	"context"

	"jobs_scanner/internal/model"
)

type Fetcher interface {
	Fetch(ctx context.Context) ([]model.Vacancy, error)
}
