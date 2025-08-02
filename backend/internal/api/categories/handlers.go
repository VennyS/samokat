package categories

import (
	"encoding/json"
	"net/http"

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
			logger.Error("Failed to get categories by warehouse ID", zap.Error(err))
			http.Error(w, "Failed to get categories", http.StatusInternalServerError)
			return
		}

		if len(categories) == 0 {
			logger.Info("No categories found for warehouse ID", zap.String("warehouse_id", warehouseID))
			http.Error(w, "No categories found", http.StatusNotFound)
			return
		}

		logger.Info("Successfully fetched categories", zap.Int("count", len(categories)))
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(categories); err != nil {
			logger.Error("Failed to encode categories to JSON", zap.Error(err))
			http.Error(w, "Failed to encode categories", http.StatusInternalServerError)
			return
		}
		logger.Info("Categories successfully returned", zap.Int("count", len(categories)))
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Categories successfully returned"))
		logger.Info("Response sent successfully")
	}
}
