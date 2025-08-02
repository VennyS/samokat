package categories

import (
	"context"
	"samokat/internal/shared"
	"samokat/internal/shared/dto"
	"samokat/internal/storage"
	"samokat/internal/storage/repository"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type CategoriesService interface {
	GetAllByWareHouseID(ctx context.Context, warehouseID uuid.UUID) ([]*dto.CategoryDTO, error)
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

func (s *categoriesSrv) GetAllByWareHouseID(ctx context.Context, warehouseID uuid.UUID) ([]*dto.CategoryDTO, error) {
	cats, err := s.categoryRepo.GetAllByWareHouseID(ctx, warehouseID)
	if err != nil {
		return nil, shared.InternalError
	}
	return mapCategoriesToDTO(cats), nil
}

func mapCategoriesToDTO(cats []*storage.Category) []*dto.CategoryDTO {
	byID := make(map[uuid.UUID]*dto.CategoryDTO)
	var roots []*dto.CategoryDTO

	for _, c := range cats {
		dtoCat := &dto.CategoryDTO{
			ID:       c.ID,
			Name:     c.Name,
			ImageURL: c.ImageURL,
		}
		byID[c.ID] = dtoCat
	}

	for _, c := range cats {
		dtoCat := byID[c.ID]
		if c.ParentID != nil {
			if parent, ok := byID[*c.ParentID]; ok {
				parent.Children = append(parent.Children, dtoCat)
			}
		} else {
			roots = append(roots, dtoCat)
		}
	}

	return roots
}
