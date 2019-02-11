package core

import (
	"github.com/dulev/ganki/server/models"
	"github.com/jinzhu/gorm"
)

type DeckService interface {
	List(username string) ([]models.Deck, error)
	Create(username string, deck models.Deck) (string, error)
	Share(username, deckID string) error
	View(username, deckID string) (models.Deck, error)
}

func NewDeckService(database *gorm.DB) DeckService {
	return &DeckServiceImpl{}
}

type DeckServiceImpl struct {
	database *gorm.DB
}

func (ds *DeckServiceImpl) List(username string) ([]models.Deck, error) {
	var decks []models.Deck
	ds.database.Where("username == ?", username).Find(&decks)
}

func (ds *DeckServiceImpl) Create(username string, deck models.Deck) (string, error) {
	panic("implement me")
}

func (ds *DeckServiceImpl) Share(username, deckID string) error {
	panic("implement me")
}

func (ds *DeckServiceImpl) View(username, deckID string) (models.Deck, error) {
	panic("implement me")
}
