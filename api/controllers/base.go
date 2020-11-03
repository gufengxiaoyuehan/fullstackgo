package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/gorilla/mux"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/gufengxiaoyuehan/fullstackgo/api/models"
)

type Server struct {
	DB *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string){
	var err error

	if Dbdriver == "mysql"{
		DBURL := fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", 
			DbUser, DbPassword, DbHost, DbPort, DbName,
		)
		server.DB, err = gorm.Open(Dbdriver, DBURL)
		if err != nil {
			fmt.Printf("Cannot connect to %s database", Dbdriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to ")
		}
	}

	if Dbdriver == "postgres"{
		DBURL := fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
			DbHost, DbPort, DbUser, DbName, DbPassword,
		)
		server.DB, err = gorm.Open(Dbdriver, DBURL)
		if err != nil {
			fmt.Printf("Cannot connect to %s database", Dbdriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to ")
		}
	}

	server.DB.Debug().AutoMigrate(&models.User{}, &models.Post{}) // database migration
	server.Router = mux.NewRouter()
	server.InitializeRoutes()
}


func (server *Server) Run(addr string){
	fmt.Println("listening to port 8080")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}


