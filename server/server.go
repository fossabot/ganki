package main

import (
	"fmt"
	"github.com/dulev/ganki/server/controllers"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/wader/gormstore"
	"net/http"
	"os"
	"time"

	"github.com/dulev/ganki/server/models"
)

type GankiServer struct {
	database       *gorm.DB
	userController controllers.UserController
}

func (gs *GankiServer) Run() {
	defer gs.database.Close()

	// Migrations
	gs.database.AutoMigrate(models.User{})

	// TODO: Create server
	r := mux.NewRouter()
	r.HandleFunc("/user/register", gs.userController.Register).Methods("POST")
	r.HandleFunc("/user/login", gs.userController.Login).Methods("POST")
	r.HandleFunc("/user/logout", gs.userController.Logout).Methods("GET")
	r.HandleFunc("/user", gs.userController.ViewInfo).Methods("GET")

	// r.HandleFunc("/deck", )
	// POST   /deck
	// GET    /deck
	// PUT    /deck
	// DELETE /deck
	// GET    /deck/deck_id
	// POST   /deck/share
	// POST   /deck/deck_id/card
	// PUT    /deck/deck_id/card/card_id
	// DELETE /deck/deck_id/card/card_id
	// GET    /deck/deck_id/study
	// POST   /deck/deck_id/study

	http.Handle("/", r)
	err := http.ListenAndServe(":8080", nil)
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}

func main() {
	gankiServer := InitializeServer()
	gankiServer.Run()
}

func NewGankiServer(
	database *gorm.DB,
	userController controllers.UserController) *GankiServer {

	return &GankiServer{
		database:       database,
		userController: userController,
	}
}

func NewDatabase() *gorm.DB {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	// TODO: How to close
	// defer db.Close()

	return db
}

func NewGormstore(database *gorm.DB) *gormstore.Store {
	// initialize and setup cleanup
	store := gormstore.New(database, []byte("secret"))
	// db cleanup every hour
	// close quit channel to stop cleanup
	quit := make(chan struct{})
	go store.PeriodicCleanup(1*time.Hour, quit)

	return store
}
