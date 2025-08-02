package categories

import (
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

type CategoriesController struct {
	logger          *zap.SugaredLogger
	categoryService CategoriesService
}

func NewController(logger *zap.SugaredLogger, categoryService CategoriesService) *CategoriesController {
	return &CategoriesController{
		logger:          logger,
		categoryService: categoryService,
	}
}

func (c CategoriesController) RegisterRoutes(r *chi.Mux) {
	r.Route("/v1/categories", func(r chi.Router) {
		r.Get("/{id}", c.GetAllByWareHouseIDHandler())
	})
}
