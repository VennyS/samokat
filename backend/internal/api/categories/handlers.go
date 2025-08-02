package categories

import (
	"net/http"
	ht "samokat/internal/lib/http"

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
			ht.SendMessage(w, r, "Failed to get categories", http.StatusInternalServerError)
			return
		}

		if len(categories) == 0 {
			logger.Info("No categories found for warehouse ID", zap.String("warehouse_id", warehouseID))
			ht.SendMessage(w, r, "No categories found", http.StatusNotFound)
			return
		}

		ht.SendJSON(w, r, map[string]any{"categories": categories}, http.StatusOK)
	}
}
