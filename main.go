package main

import (
	// "encoding/json"
	// "io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/uptrace/bunrouter"
	// "gorm.io/driver/mysql"
	// "gorm.io/gorm"
)

type User struct {
	ID             int        `json:"id"`
	Email          string     `json:"email"`
	Name           string     `json:"name"`
	PhoneNumber    string     `json:"phone_number"`
	Address        string     `json:"address"`
	Password       string     `json:"password"`
	IsUserActive   bool       `json:"is_user_active"`
	VerificationAt time.Time  `json:"verification_at"`
	ProfileImage   string     `json:"profile_image"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
	DeletedAt      *time.Time `json:"deleted_at"`
}

func main() {
	// routes
	r := bunrouter.New()
	r.GET("/", func(w http.ResponseWriter, req bunrouter.Request) error {
		bunrouter.JSON(w, bunrouter.H{
			"message": "pong",
		})
		return nil
	})

	port := ":5000"
	log.Printf("running on port %v", port)
	log.Println(http.ListenAndServe(port, r))
}
