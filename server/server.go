package main

import (
	"fmt"
	"github.com/wader/gormstore"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/dulev/ganki/server/models"
	"github.com/dulev/ganki/server/user"
)

type GankiServer struct {
	userController user.UserController
}

func NewGankiServer(
	userController user.UserController) *GankiServer {
	return &GankiServer{
		userController: userController,
	}
}



func main() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// initialize and setup cleanup
	store := gormstore.New(db, []byte("secret"))
	// db cleanup every hour
	// close quit channel to stop cleanup
	quit := make(chan struct{})
	go store.PeriodicCleanup(1*time.Hour, quit)

	// Migrations
	db.AutoMigrate(models.User{})

	// TODO: Create server
	userController := user.NewUserController(db, user.NewUserService(db))

	r := mux.NewRouter()
	r.HandleFunc("/user/register", userController.Register)
	r.HandleFunc("/user/login", userController.Login)

	// Middleware
	// r.Use(func(next http.Handler) http.Handler {
	// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//
	// 	})
	// })

	// POST   /user/register
	// POST   /user/login
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
	err = http.ListenAndServe(":8080", nil)
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
