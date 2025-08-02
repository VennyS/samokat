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

type CategoriesSrv struct {
	logger       *zap.Logger
	categoryRepo repository.CategoryRepository
}

func NewService(logger *zap.Logger, categoryRepo repository.CategoryRepository) *CategoriesSrv {
	return &CategoriesSrv{
		logger:       logger,
		categoryRepo: categoryRepo,
	}
}

func (s *CategoriesSrv) GetAllByWareHouseID(ctx context.Context, warehouseID uuid.UUID) ([]*storage.Category, error) {
	return s.categoryRepo.GetAllByWareHouseID(ctx, warehouseID)
}
