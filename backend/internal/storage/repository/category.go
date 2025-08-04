package repository

import (
	"context"
	"samokat/internal/shared/dto"
	"samokat/internal/storage"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type CategoryRepository interface {
	GetAllByWareHouseID(ctx context.Context, warehouseID uuid.UUID) ([]*storage.Category, error)
	Create(ctx context.Context, category *dto.CreateCategoryDTO) (*storage.Category, error)
}

type categoryRepo struct {
	logger *zap.SugaredLogger
	db     *sqlx.DB
}

func NewCategoryRepo(logger *zap.SugaredLogger, db *sqlx.DB) *categoryRepo {
	return &categoryRepo{
		logger: logger,
		db:     db,
	}
}

func (r *categoryRepo) GetAllByWareHouseID(ctx context.Context, warehouseID uuid.UUID) ([]*storage.Category, error) {
	logger := r.logger.With(
		zap.String("method", "repository.GetAllByWareHouseID"),
		zap.String("warehouse_id", warehouseID.String()),
	)
	logger.Info("Fetching all categories by warehouse ID")

	var categories []*storage.Category
	query := `
		SELECT c.*
		FROM categories c
		JOIN warehouse_categories wc ON c.id = wc.category_id
		WHERE wc.warehouse_id = $1
		ORDER BY c.name
	`

	if err := r.db.SelectContext(ctx, &categories, query, warehouseID); err != nil {
		r.logger.Errorf("Failed to get categories by warehouse ID: %v", err)
		return nil, err
	}
	return categories, nil
}

func (r *categoryRepo) Create(ctx context.Context, category *dto.CreateCategoryDTO) (*storage.Category, error) {
	logger := r.logger.With(
		zap.String("method", "repository.Create"),
		zap.String("category_name", category.Name),
	)
	logger.Info("Creating new category")
	query := `
		INSERT INTO categories (name, image_url, parent_id)
		VALUES ($1, $2, $3) RETURNING id, name, image_url, parent_id
	`
	var newCategory storage.Category
	err := r.db.QueryRowContext(ctx, query, category.Name, category.ImageURL, category.ParentID).Scan(&newCategory.ID, &newCategory.Name, &newCategory.ImageURL, &newCategory.ParentID)
	if err != nil {
		logger.Errorf("Failed to create category: %v", err)
		return nil, err
	}
	return &newCategory, nil
}
