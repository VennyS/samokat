package categories

import (
	"context"
	"errors"
	"samokat/internal/shared"
	"samokat/internal/shared/dto"
	"samokat/internal/storage"
	"samokat/internal/storage/repository"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type CategoriesService interface {
	GetAllByWareHouseID(ctx context.Context, warehouseID uuid.UUID) ([]*dto.CategoryDTO, error)
	Create(ctx context.Context, category *dto.CreateCategoryDTO) (*dto.CategoryDTO, error)
	Delete(ctx context.Context, id uuid.UUID) error
	Put(ctx context.Context, categoryID uuid.UUID, category *dto.UpdateCategoryDTO) (*dto.CategoryDTO, error)
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
	logger := s.logger.With(
		zap.String("method", "service.GetAllByWareHouseID"),
		zap.String("warehouse_id", warehouseID.String()),
	)
	logger.Info("Fetching all categories by warehouse ID")

	cats, err := s.categoryRepo.GetAllByWareHouseID(ctx, warehouseID)
	if err != nil {
		return nil, shared.InternalError
	}

	return mapCategoriesToDTO(cats), nil
}

func (s *categoriesSrv) Create(ctx context.Context, category *dto.CreateCategoryDTO) (*dto.CategoryDTO, error) {
	logger := s.logger.With(
		zap.String("method", "service.Create"),
		zap.String("category_name", category.Name),
	)
	logger.Info("Creating new category")

	newCategory, err := s.categoryRepo.Create(ctx, category)
	if err != nil {
		logger.Errorf("Failed to create category: %v", err)
		return nil, shared.InternalError
	}
	return mapCategoryToDTO(newCategory), nil
}

func (s *categoriesSrv) Delete(ctx context.Context, id uuid.UUID) error {
	logger := s.logger.With(
		zap.String("method", "service.Delete"),
		zap.String("category_id", id.String()),
	)
	logger.Info("Deleting category")

	if err := s.categoryRepo.Delete(ctx, id); err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			logger.Warnf("Category not found for deletion: %v", err)
			return shared.NotFoundError
		}

		logger.Errorf("Failed to delete category: %v", err)
		return shared.InternalError
	}
	return nil
}

func (s *categoriesSrv) Put(ctx context.Context, categoryID uuid.UUID, category *dto.UpdateCategoryDTO) (*dto.CategoryDTO, error) {
	logger := s.logger.With(
		zap.String("method", "service.Update"),
		zap.String("category_id", categoryID.String()),
	)
	logger.Info("Updating category")

	updatedCategory, err := s.categoryRepo.Put(ctx, categoryID, category)
	if err != nil {
		logger.Errorf("Failed to update category: %v", err)
		return nil, shared.InternalError
	}
	return mapCategoryToDTO(updatedCategory), nil
}

// mapCategoryToDTO maps a storage.Category to a dto.CategoryDTO.
func mapCategoryToDTO(c *storage.Category) *dto.CategoryDTO {
	if c == nil {
		return nil
	}
	return &dto.CategoryDTO{
		ID:       c.ID,
		Name:     c.Name,
		ImageURL: c.ImageURL,
		// Children is omitted here; typically used in tree structures.
	}
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
