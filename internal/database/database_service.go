package database

import (
	"context"

	"github.com/cwrenhold/linkedin-learning-build-a-microservice-with-go/internal/models"
)

func (c Client) GetAllServices(ctx context.Context) ([]models.Service, error) {
	var services []models.Service
	result := c.DB.WithContext(ctx).
		Where(models.Service{}).
		Find(&services)

	return services, result.Error
}
