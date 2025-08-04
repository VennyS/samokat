package repository

import (
	"context"
	"samokat/internal/storage"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type CategoryRepository interface {
	GetAllByWareHouseID(ctx context.Context, warehouseID uuid.UUID) ([]*storage.Category, error)
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
