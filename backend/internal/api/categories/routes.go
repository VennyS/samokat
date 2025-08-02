package categories

import (
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

type CategoriesController struct {
	logger          *zap.Logger
	categoryService CategoriesService
}

func NewController(logger *zap.Logger, categoryService CategoriesService) *CategoriesController {
	return &CategoriesController{
		logger:          logger,
		categoryService: categoryService,
	}
}

func (c CategoriesController) RegisterRoutes(r *chi.Mux) {
	r.Route("v1/categories", func(r chi.Router) {
		r.Get("/{id}", c.GetAllByWareHouseIDHandler())
	})
}
