package main

import (
	"net/http"

	"github.com/b1994mi/golang-rest-api-example/handler/auth"
	"github.com/b1994mi/golang-rest-api-example/handler/transaction"
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
	userTokenRepo := model.NewUserTokenRepo(db)
	userTransationRepo := model.NewUserTransactionRepo(db)

	// routes with handlers
	userHandler := user.NewHandler(
		userRepo,
	)

	routes.POST("/register", util.MakeHandler(
		userHandler.CreateHandler,
		util.ShouldBindJSON,
	))

	routes.GET("/user/:user_id", util.MakeHandler(
		userHandler.FindHandler,
		util.ShouldBindUri,
	))

	authHandler := auth.NewHandler(
		userRepo,
		userTokenRepo,
	)

	routes.POST("/login", util.MakeHandler(
		authHandler.LoginHandler,
		util.ShouldBindJSON,
	))

	routes.POST("/refresh-token", util.MakeHandler(
		authHandler.RefreshTokenHandler,
		util.ShouldBindJSON,
	))

	transactionHandler := transaction.NewHandler(
		userRepo,
		userTransationRepo,
	)

	routes.POST("/topup", util.MakeHandler(
		transactionHandler.TopUpHandler,
		util.ShouldBindJWT,
		util.ShouldBindJSON,
	))

	return routes
}
