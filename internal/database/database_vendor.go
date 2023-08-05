package database

import (
	"context"

	"github.com/cwrenhold/linkedin-learning-build-a-microservice-with-go/internal/models"
)

func (c Client) GetAllVendors(ctx context.Context) ([]models.Vendor, error) {
	var vendors []models.Vendor
	result := c.DB.WithContext(ctx).
		Where(models.Vendor{}).
		Find(&vendors)

	return vendors, result.Error
}
