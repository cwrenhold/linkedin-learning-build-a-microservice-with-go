package database

import (
	"context"

	"github.com/cwrenhold/linkedin-learning-build-a-microservice-with-go/internal/models"
)

func (c Client) GetAllProducts(ctx context.Context, vendorID string) ([]models.Product, error) {
	var products []models.Product
	result := c.DB.WithContext(ctx).
		Where(models.Product{VendorID: vendorID}).
		Find(&products)

	return products, result.Error
}
