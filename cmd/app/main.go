package main

import (
	"fmt"
	"github.com/anisbouzahar/portfolio-api/docs"
	"github.com/anisbouzahar/portfolio-api/pkg/helpers"
	"github.com/spf13/pflag"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

var (
	version string
	addr    string
)

func init() {
	pflag.StringVarP(&addr, "address", "a", ":8080", "the address for the api to listen on. Host and port separated by ':'")
	pflag.Parse()
}

func main() {
	docs.SwaggerInfo.Version = version
	docs.SwaggerInfo.BasePath = addr
	docs.SwaggerInfo.BasePath = helpers.GetEnvWithDefault("API_BASE_PATH", "/v1")
	docs.SwaggerInfo.Schemes = strings.Split(helpers.GetEnvWithDefault("API_PROTOCOL", "http"), ",")

	// gracefully exit on keyboard interrupt
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	server, cleanup, err := InitializeServer()
	if err != nil {
		fmt.Printf("unable to start application %v", err)
	}
	logger := server.Api.Logger
	defer cleanup()

	server.Api.SetUpRoutes()

	// print current version
	server.Api.Logger.Info("starting up API...", zap.String("version", version))

	go func() {
		if err := http.ListenAndServe(addr, server.Api.R); err != nil {
			logger.Error("failed to start server", zap.Error(err))
			os.Exit(1)
		}
	}()
	logger.Info("ready to serve requests on " + addr)

	<-c
	logger.Info("gracefully shutting down")
	os.Exit(0)
}
