package middleware

import (
	"github.com/dulev/ganki/server/common"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
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

func (sm *SessionManager) Login(username string, w http.ResponseWriter, r *http.Request) {
	session, err := sm.sessionStore.Get(r, SessionCookieName)
	if err != nil {
		panic("couldn't get session")
	}

	session.Values["username"] = username

	err = session.Save(r, w)
	if err != nil {
		panic("hanle error")
	}
}

func (sm *SessionManager) ShouldBeLoggedIn(w http.ResponseWriter, r *http.Request) (string, error) {
	session, err := sm.sessionStore.Get(r, common.SessionStoreName)
	if err != nil {
		panic("TODO")
	}

	usernameAsInterface, exists := session.Values["username"]
	if !exists {
		w.WriteHeader(http.StatusUnauthorized)
	}

	username, err := usernameAsInterface.(string)
	if err != nil {
		return "", errors.Wrap(err, "TODO")
	}

	return username, nil
}

func (sm *SessionManager) ShouldBeLoggedOut(w http.ResponseWriter, r *http.Request) {
	session, err := sm.sessionStore.Get(r, common.SessionStoreName)
	if err != nil {
		panic("TODO")
	}

	_, exists := session.Values["username"]
	if exists {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func (sm *SessionManager) Logout(w http.ResponseWriter, r *http.Request) {
	session, err := sm.sessionStore.Get(r, common.SessionStoreName)
	if err != nil {
		panic("TODO")
	}

	session.Values = make(map[interface{}]interface{})
	err = session.Save(r, w)
	if err != nil {
		panic("TODO alabala")
	}
}
