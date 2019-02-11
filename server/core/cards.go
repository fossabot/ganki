package core

import "github.com/dulev/ganki/server/models"

type CardService interface {
	AddCard(username, deckID string, card models.Card) (string, error)
	RemoveCard(username, deckID string, cardID string) error
	UpdateCard(username, deckID string, card models.Card) error
}
