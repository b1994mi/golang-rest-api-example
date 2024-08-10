package main

import (
	"net/http"

	"github.com/b1994mi/golang-rest-api-example/handler/user"
	"github.com/b1994mi/golang-rest-api-example/model"
	"github.com/b1994mi/golang-rest-api-example/util"
	"github.com/uptrace/bunrouter"
	"gorm.io/gorm"
)

func setupRoutes(
	db *gorm.DB,
) *bunrouter.Router {
	routes := bunrouter.New()
	routes.GET("/", func(w http.ResponseWriter, req bunrouter.Request) error {
		bunrouter.JSON(w, bunrouter.H{
			"message": "pong",
		})
		return nil
	})

	// init all repos for dependency injection
	userRepo := model.NewUserRepo(db)

	// routes with handlers
	userHandler := user.NewHandler(
		userRepo,
	)

	routes.GET("/user", util.MakeHandler(userHandler.FindHandler))
	routes.POST("/user", util.MakeHandler(userHandler.CreateHandler))
	routes.POST("/verify", util.MakeHandler(userHandler.VerifyHandler))

	return routes
}