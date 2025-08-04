package repository

import (
	"context"
	"database/sql"
	"errors"
	"samokat/internal/shared/dto"
	"samokat/internal/storage"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

var (
	ErrNotFound       = errors.New("not found")
	ErrParentNotFound = errors.New("parent did not exists")
)

type CategoryRepository interface {
	GetAllByWareHouseID(ctx context.Context, warehouseID uuid.UUID) ([]*storage.Category, error)
	Create(ctx context.Context, category *dto.CreateCategoryDTO) (*storage.Category, error)
	Delete(ctx context.Context, id uuid.UUID) error
	Put(ctx context.Context, categoryID uuid.UUID, category *dto.UpdateCategoryDTO) (*storage.Category, error)
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

	if category.ParentID != nil {
		var exists bool
		err := r.db.GetContext(ctx, &exists, `SELECT EXISTS(SELECT 1 FROM categories WHERE id = $1)`, *category.ParentID)
		if err != nil {
			logger.Errorf("Failed to check parent category: %v", err)
			return nil, err
		}
		if !exists {
			return nil, ErrParentNotFound
		}
	}

	query := `
		INSERT INTO categories (name, image_url, parent_id)
		VALUES ($1, $2, $3)
		RETURNING id, name, image_url, parent_id
	`
	var newCategory storage.Category
	err := r.db.QueryRowContext(ctx, query, category.Name, category.ImageURL, category.ParentID).
		Scan(&newCategory.ID, &newCategory.Name, &newCategory.ImageURL, &newCategory.ParentID)
	if err != nil {
		logger.Errorf("Failed to create category: %v", err)
		return nil, err
	}
	return &newCategory, nil
}

func (r *categoryRepo) Delete(ctx context.Context, id uuid.UUID) error {
	logger := r.logger.With(
		zap.String("method", "repository.Delete"),
		zap.String("category_id", id.String()),
	)
	logger.Info("Deleting category")

	query := `
		DELETE FROM categories
		WHERE id = $1
	`

	if _, err := r.db.ExecContext(ctx, query, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			logger.Warnf("Category not found for deletion: %v", err)
			return ErrNotFound
		}

		logger.Errorf("Failed to delete category: %v", err)
		return err
	}
	return nil
}

func (r *categoryRepo) Put(ctx context.Context, categoryID uuid.UUID, category *dto.UpdateCategoryDTO) (*storage.Category, error) {
	logger := r.logger.With(
		zap.String("method", "repository.Put"),
		zap.String("category_name", *category.Name),
	)
	logger.Info("Updating category")

	if category.ParentID != nil {
		var exists bool
		err := r.db.GetContext(ctx, &exists, `SELECT EXISTS(SELECT 1 FROM categories WHERE id = $1)`, *category.ParentID)
		if err != nil {
			logger.Errorf("Failed to check parent category: %v", err)
			return nil, err
		}
		if !exists {
			return nil, ErrParentNotFound
		}
	}

	var updatedCategory storage.Category
	query := `
		UPDATE categories
		SET name = COALESCE($1, name),
			image_url = COALESCE($2, image_url),
			parent_id = COALESCE($3, parent_id)
		WHERE id = $4
		RETURNING id, name, image_url, parent_id
	`
	err := r.db.QueryRowContext(ctx, query, category.Name, category.ImageURL, category.ParentID, categoryID).Scan(&updatedCategory.ID, &updatedCategory.Name, &updatedCategory.ImageURL, &updatedCategory.ParentID)
	if err != nil {
		logger.Errorf("Failed to update category: %v", err)
		return nil, err
	}
	return &updatedCategory, nil
}
