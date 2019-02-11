// +build wireinject

package main

import (
	"github.com/dulev/ganki/server/controllers"
	"github.com/dulev/ganki/server/controllers/middleware"
	"github.com/dulev/ganki/server/user"
)
import "github.com/google/wire"

func InitializeServer() *GankiServer {
	wire.Build(
		NewGankiServer,
		NewDatabase,

		// Middleware
		middleware.NewSessionManager,

		// Controllers
		controllers.NewUserController,

		// Services
		user.NewUserService,
	)
	return &GankiServer{}
}
