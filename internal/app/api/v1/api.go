package v1

import (
	"github.com/anisbouzahar/portfolio-api/internal/app/handlers/user"
	m "github.com/anisbouzahar/portfolio-api/internal/app/middlewares"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger"
	"go.uber.org/zap"
	"net/http"
)

type API struct {
	R           *chi.Mux
	Logger      *zap.Logger
	userHandler *user.Handler
}

func NewAPI(logger *zap.Logger, r *chi.Mux, userH *user.Handler) *API {
	r.Use(middleware.RequestID)

	if logger != nil {
		r.Use(m.SetLogger(logger))
	}

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
	}))
	return &API{
		Logger:      logger,
		R:           r,
		userHandler: userH,
	}
}

func (api *API) SetUpRoutes() {
	api.R.HandleFunc("/swagger", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, r.RequestURI+"/", http.StatusMovedPermanently)
	})
	api.R.Get("/swagger*", httpSwagger.Handler())

	api.R.Route("/v1", func(r chi.Router) {
		r.Post("/subscribe", api.userHandler.SubscribeTobeNotified)
	})
}
