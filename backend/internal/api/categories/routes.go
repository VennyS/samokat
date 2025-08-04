package categories

import (
	"samokat/internal/api/middleware"
	"samokat/internal/shared/dto"

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
		r.With(middleware.JsonBodyMiddleware[*dto.CreateCategoryDTO](c.logger)).Post("/", (c.CreateHandler()))
		r.Delete("/{id}", c.DeleteHandler())
		r.With(middleware.JsonBodyMiddleware[*dto.UpdateCategoryDTO](c.logger)).Put("/{id}", c.PutHandler())
	})
}
