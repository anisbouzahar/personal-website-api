// wire.go

//go:build wireinject
// +build wireinject

package main

import (
	"github.com/anisbouzahar/portfolio-api/internal/app/api/v1"
	"github.com/anisbouzahar/portfolio-api/internal/app/database"
	userHandler "github.com/anisbouzahar/portfolio-api/internal/app/handlers/user"
	"github.com/anisbouzahar/portfolio-api/internal/app/http"
	userRepository "github.com/anisbouzahar/portfolio-api/internal/app/repository/user"
	"github.com/anisbouzahar/portfolio-api/internal/app/services/user"
	"github.com/anisbouzahar/portfolio-api/pkg"
	"github.com/go-chi/chi/v5"
	"github.com/google/wire"
)

func InitializeServer() (*http.Server, func(), error) {
	wire.Build(http.NewServer, pkg.NewLogger, v1.NewAPI, chi.NewRouter, userHandler.NewHandler, user.NewUserService, userRepository.NewUserRepository, database.NewMongoDb)
	return &http.Server{}, func() {}, nil
}
