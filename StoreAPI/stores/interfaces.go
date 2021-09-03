package stores

import (
	"context"

	"github.com/AmazonScraper/StoreAPI/models"
)

type Store interface {
	Create(ctx context.Context, product models.ProductDocument) error
}
