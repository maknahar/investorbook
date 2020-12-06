package routes

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/raksul-code-review/userapi-candidate-maknahar-a993286a1d8d72e3a9534ec66ef11449/internal/configs"
)

func Get(conf *configs.Conf) *chi.Mux {
	r := chi.NewRouter()

	cor := cors.New(cors.Options{
		AllowedOrigins:     []string{"*"}, // Change this according to url of env
		AllowedMethods:     []string{http.MethodGet, http.MethodOptions},
		AllowedHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:     []string{"Link"},
		AllowCredentials:   true,
		OptionsPassthrough: true,
		MaxAge:             300, // Maximum value not ignored by any of major browsers
	})

	// Set a timeout value on the request context (ctx), that will signal through ctx.Done() that the request has timed
	// out and further processing should be stopped.
	r.Use(middleware.Timeout(time.Minute), middleware.AllowContentType("application/json"))

	r.Use(cor.Handler, middleware.RequestID, middleware.RealIP, middleware.Logger, middleware.Recoverer)

	user := NewReportHandler(conf.DB)

	r.Get("/companies/{companyId}/export", user.GenerateCompanyReport)

	return r
}
