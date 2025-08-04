package categories

import (
	"net/http"
	"samokat/internal/api/middleware"
	ht "samokat/internal/lib/http"
	"samokat/internal/shared/dto"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func (c CategoriesController) GetAllByWareHouseIDHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		warehouseID := chi.URLParam(r, "id")
		logger := c.logger.With(
			zap.String("method", "categoriesController.GetAllByWareHouseIDHandler"),
			zap.String("warehouse_id", warehouseID),
		)

		categories, err := c.categoryService.GetAllByWareHouseID(r.Context(), uuid.MustParse(warehouseID))
		if err != nil {
			logger.Errorf("Failed to get categories by warehouse ID: %v", err)
			ht.SendMessage(w, r, "Failed to get categories", http.StatusInternalServerError)
			return
		}

		if len(categories) == 0 {
			logger.Infof("No categories found for warehouse ID: %s", warehouseID)
			ht.SendMessage(w, r, "No categories found", http.StatusNotFound)
			return
		}

		ht.SendJSON(w, r, map[string]any{"categories": categories}, http.StatusOK)
	}
}

func (c CategoriesController) CreateHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := r.Context().Value(middleware.DataKey).(*dto.CreateCategoryDTO)

		logger := c.logger.With(
			zap.String("method", "categoriesController.CreateHandler"),
			zap.String("category_name", req.Name),
		)

		newCategory, err := c.categoryService.Create(r.Context(), req)
		if err != nil {
			logger.Errorf("Failed to create category: %v", err)
			ht.SendMessage(w, r, "Failed to create category", http.StatusInternalServerError)
			return
		}

		ht.SendJSON(w, r, map[string]any{"category": newCategory}, http.StatusCreated)
	}
}

func (c CategoriesController) DeleteHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		idStr := chi.URLParam(r, "id")
		logger := c.logger.With(
			zap.String("method", "categoriesController.DeleteHandler"),
			zap.String("category_id", idStr),
		)

		id, err := uuid.Parse(idStr)
		if err != nil {
			logger.Errorf("Invalid category ID: %v", err)
			ht.SendMessage(w, r, "Invalid category ID", http.StatusBadRequest)
			return
		}

		if err := c.categoryService.Delete(r.Context(), id); err != nil {
			logger.Errorf("Failed to delete category: %v", err)
			ht.SendMessage(w, r, "Failed to delete category", http.StatusInternalServerError)
			return
		}

		ht.SendMessage(w, r, "Category deleted successfully", http.StatusOK)
	}
}

func (c CategoriesController) PutHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		updateReq := r.Context().Value(middleware.DataKey).(*dto.UpdateCategoryDTO)
		idStr := chi.URLParam(r, "id")
		logger := c.logger.With(
			zap.String("method", "categoriesController.PutHandler"),
			zap.String("category_id", idStr),
		)

		if updateReq.Name == nil && updateReq.ImageURL == nil && updateReq.ParentID == nil {
			ht.SendMessage(w, r, "No fields to update", http.StatusBadRequest)
			return
		}

		id, err := uuid.Parse(idStr)
		if err != nil {
			logger.Errorf("Invalid category ID: %v", err)
			ht.SendMessage(w, r, "Invalid category ID", http.StatusBadRequest)
			return
		}

		updatedCategory, err := c.categoryService.Put(r.Context(), id, updateReq)
		if err != nil {
			logger.Errorf("Failed to update category: %v", err)
			ht.SendMessage(w, r, "Failed to update category", http.StatusInternalServerError)
			return
		}

		ht.SendJSON(w, r, map[string]any{"category": updatedCategory}, http.StatusOK)
	}
}
