package src

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/keloran/go-probe"
	"github.com/keloran/go-healthcheck"
	"os"
	"time"
)

// Routes router
func Routes() chi.Router {
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(60 * time.Second))

	// Probe
	router.Get("/probe", probe.HTTP)
	router.Get(fmt.Sprintf("%s/probe", os.Getenv("SITE_PREFIX")), probe.HTTP)

	// HealthCheck
	router.Get(fmt.Sprintf("%s/healthcheck", os.Getenv("SITE_PREFIX")), healthcheck.HTTP)

	return router
}