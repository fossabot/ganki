package core

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

type StudySessionManager interface {
	IncreaseDailyNewCards(userID, deckID string, increase int) error
	UpdateCardsPerDay(userID, deckID string, cpd int) error

	Study(userID, deckID string) Card
	SubmitStudyResults(userID, deckID string) error
}
