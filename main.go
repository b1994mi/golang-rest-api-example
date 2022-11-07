package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/b1994mi/golang-rest-api-example/model"
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

	r.GET("/user", func(w http.ResponseWriter, req bunrouter.Request) error {
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			bunrouter.JSON(w, bunrouter.H{"message": err})
			return nil
		}

		var reqBody struct {
			ID int `json:"id"`
		}
		err = json.Unmarshal(body, &reqBody)
		if err != nil {
			bunrouter.JSON(w, bunrouter.H{"message": err})
			return nil
		}

		var m model.User

		err = db.Where("id", reqBody.ID).Find(&m).Error
		if err != nil {
			bunrouter.JSON(w, bunrouter.H{
				"message": err,
			})
		}

		bunrouter.JSON(w, bunrouter.H{
			"data": m,
		})
		return nil
	})

	r.POST("/user", func(w http.ResponseWriter, req bunrouter.Request) error {
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			bunrouter.JSON(w, bunrouter.H{"message": err})
			return nil
		}

		var reqBody struct {
			Name string `json:"name"`
		}
		err = json.Unmarshal(body, &reqBody)
		if err != nil {
			bunrouter.JSON(w, bunrouter.H{"message": err})
			return nil
		}
		now := time.Now()
		err = db.Create(&model.User{
			Email:          "",
			Name:           reqBody.Name,
			PhoneNumber:    "",
			Address:        "",
			Password:       "",
			IsUserActive:   false,
			VerificationAt: now,
			ProfileImage:   "",
			CreatedAt:      now,
		}).Error
		if err != nil {
			bunrouter.JSON(w, bunrouter.H{
				"message": err,
			})
		}

		bunrouter.JSON(w, bunrouter.H{
			"acknowledge": true,
		})
		return nil
	})

	r.POST("/verify", func(w http.ResponseWriter, req bunrouter.Request) error {
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			bunrouter.JSON(w, bunrouter.H{"message": err})
			return nil
		}

		var reqBody struct {
			ID int `json:"id"`
		}
		err = json.Unmarshal(body, &reqBody)
		if err != nil {
			bunrouter.JSON(w, bunrouter.H{"message": err})
			return nil
		}

		var m model.User

		err = db.Where("id", reqBody.ID).Find(&m).Error
		if err != nil {
			bunrouter.JSON(w, bunrouter.H{
				"message": err,
			})
		}

		m.IsUserActive = true

		err = db.Save(&m).Debug().Error
		if err != nil {
			bunrouter.JSON(w, bunrouter.H{
				"message": err,
			})
		}

		bunrouter.JSON(w, bunrouter.H{
			"acknowledge": true,
		})
		return nil
	})

	port := ":5000"
	log.Printf("running on port %v", port)
	log.Println(http.ListenAndServe(port, r))
}
