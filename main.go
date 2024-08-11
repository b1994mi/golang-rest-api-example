package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/b1994mi/golang-rest-api-example/handler/transaction"
	"github.com/b1994mi/golang-rest-api-example/message"
	"github.com/b1994mi/golang-rest-api-example/model"
	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	dsn := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_URL"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SCHEMA"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Println("Failed Initializing Broker Connection")
		panic(err)
	}

	routes := setupRoutes(db, conn)

	// TODO: tidy up this connection function, maybe setup rmq or whatever

	ch, err := conn.Channel()
	if err != nil {
		log.Println(err)
	}

	q, err := ch.QueueDeclare(
		"transfer",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}

	log.Println(q)

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
		panic(err)
	}

	go func() {
		// TODO: make this as a makeHandler style method for dependency injection
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

	port := ":5000"
	log.Printf("running on port %v", port)
	log.Println(http.ListenAndServe(port, routes))
}
