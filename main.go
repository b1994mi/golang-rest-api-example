package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/b1994mi/golang-rest-api-example/handler"
	"github.com/joho/godotenv"
	"github.com/uptrace/bunrouter"
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

	// routes
	r := bunrouter.New()
	r.GET("/", func(w http.ResponseWriter, req bunrouter.Request) error {
		bunrouter.JSON(w, bunrouter.H{
			"message": "pong",
		})
		return nil
	})

	// routes with handlers
	h := handler.NewHandler(db)
	r.GET("/user", h.FindHandler)
	r.POST("/user", h.CreateHandler)
	r.POST("/verify", h.VerifyHandler)

	port := ":5000"
	log.Printf("running on port %v", port)
	log.Println(http.ListenAndServe(port, r))
}
