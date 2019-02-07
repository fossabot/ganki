package auth

import (
	"github.com/jinzhu/gorm"
	"github.com/wader/gormstore"
	"net/http"
	"time"
)

// initialize and setup cleanup

const SessionCookieName = "testCookieName"

type SessionManager struct {
	sessionStore *gormstore.Store
}

func NewSessionManager(database *gorm.DB) *SessionManager {
	store := gormstore.New(database, []byte("dev-secret"))
	quit := make(chan struct{})
	go store.PeriodicCleanup(1*time.Hour, quit)

	return &SessionManager{sessionStore: store}
}

func (sm *SessionManager) sessionLogin(w http.ResponseWriter, r *http.Request) {
	session, err := sm.sessionStore.Get(r, SessionCookieName)
	if err != nil {
		panic("couldn't get session")
	}

	if !session.IsNew {
		panic("session shouldn't exist when logging in")
	}

	session.Values[""]
	session.Values["authenticated"] = true
}

func (sm *SessionManager) sessionAuthenticate(w http.ResponseWriter, r *http.Request) {

}

func (sm *SessionManager) sessionLogout(w http.ResponseWriter, r *http.Request) {
}
