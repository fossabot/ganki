package core

type Card struct {
	ID    string
	Front CardFace
	Back  CardFace
}

type CardFace struct {
	PrimaryInfo   string
	SecondaryInfo string
	Image         string
}

type CardManager interface {
	AddCard(userID, deckID string, card Card) (string, error)
	RemoveCard(userID, deckID string, cardID string) error
	UpdateCard(userID, deckID string, card Card) error
}
