// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"context"
	"net/http"

	"github.com/Pranc1ngPegasus/connect-go-practice/adapter/handler"
	"github.com/Pranc1ngPegasus/connect-go-practice/adapter/server"
	"github.com/Pranc1ngPegasus/connect-go-practice/infra/configuration"
	"github.com/Pranc1ngPegasus/connect-go-practice/infra/logger"
)

// Injectors from wire.go:

func initialize() (*app, error) {
	contextContext := context.Background()
	loggerLogger, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}
	configurationConfiguration, err := configuration.NewConfiguration()
	if err != nil {
		return nil, err
	}
	apiv1Handler := handler.NewAPIV1Handler(loggerLogger)
	httpServer := server.NewServer(contextContext, loggerLogger, configurationConfiguration, apiv1Handler)
	mainApp := &app{
		server: httpServer,
	}
	return mainApp, nil
}

// wire.go:

type app struct {
	server *http.Server
}
