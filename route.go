package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/b1994mi/golang-rest-api-example/handler/auth"
	"github.com/b1994mi/golang-rest-api-example/handler/transaction"
	"github.com/b1994mi/golang-rest-api-example/handler/user"
	"github.com/b1994mi/golang-rest-api-example/message"
	"github.com/b1994mi/golang-rest-api-example/model"
	"github.com/b1994mi/golang-rest-api-example/util"
	"github.com/streadway/amqp"
	"github.com/uptrace/bunrouter"
	"gorm.io/gorm"
)

func setupRoutes(
	db *gorm.DB,
	conn *amqp.Connection,
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
	userTransactionRepo := model.NewUserTransactionRepo(db)
	transferRepo := message.NewTransferRepo(conn)

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
		userTransactionRepo,
		transferRepo,
	)

	routes.POST("/topup", util.MakeHandler(
		transactionHandler.TopUpHandler,
		util.ShouldBindJWT,
		util.ShouldBindJSON,
	))

	routes.POST("/transfer", util.MakeHandler(
		transactionHandler.TransferHandler,
		util.ShouldBindJWT,
		util.ShouldBindJSON,
	))

	return routes
}

func setupConsumer(
	db *gorm.DB,
	conn *amqp.Connection,
) error {
	ch, err := conn.Channel()
	if err != nil {
		return err
	}

	_, err = ch.QueueDeclare(
		"transfer",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	msgs, err := ch.Consume(
		"transfer",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	go func() {
		// TODO: make this as a makeHandler style method for dependency injection
		// and easier to create new consumers for different queue
		userRepo := model.NewUserRepo(db)
		userTransactionRepo := model.NewUserTransactionRepo(db)
		transferRepo := message.NewTransferRepo(conn)

		transactionHandler := transaction.NewHandler(
			userRepo,
			userTransactionRepo,
			transferRepo,
		)

		var data message.Transfer
		for d := range msgs {
			log.Printf("Recieved Message: %s\n", d.Body)

			err := json.Unmarshal(d.Body, &data)
			if err != nil {
				log.Println(err)
			}

			err = transactionHandler.TransferConsumer(&data)
			if err != nil {
				log.Println(err)
			}
		}
	}()

	return nil
}
