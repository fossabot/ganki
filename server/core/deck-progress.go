package core

import "github.com/dulev/ganki/server/models"

type DeckUserData struct {
	UserID        string
	DeckID        string
	CardsPerDay   int
	DailyIncrease int
	Progress      DeckProgress
}

type DeckProgress struct {
	DailyNewCards   int
	NewCardIncrease int
	CardProgress    map[string]CardProgress
}

type CardProgress struct {
}

type StudySessionService interface {
	IncreaseDailyNewCards(userID, deckID string, increase int) error
	UpdateCardsPerDay(userID, deckID string, cpd int) error

	Study(userID, deckID string) models.Card
	SubmitStudyResults(userID, deckID string) error
}
