package categories

import (
	"context"
	"samokat/internal/storage"
	"samokat/internal/storage/repository"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type CategoriesService interface {
	GetAllByWareHouseID(ctx context.Context, warehouseID uuid.UUID) ([]*storage.Category, error)
}

type categoriesSrv struct {
	logger       *zap.SugaredLogger
	categoryRepo repository.CategoryRepository
}

func NewService(logger *zap.SugaredLogger, categoryRepo repository.CategoryRepository) *categoriesSrv {
	return &categoriesSrv{
		logger:       logger,
		categoryRepo: categoryRepo,
	}
}

func (s *categoriesSrv) GetAllByWareHouseID(ctx context.Context, warehouseID uuid.UUID) ([]*storage.Category, error) {
	return s.categoryRepo.GetAllByWareHouseID(ctx, warehouseID)
}
