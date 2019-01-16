package core

type Deck struct {
	ID    string
	Owner string
	Name  string
	Cards []Card
}

type DeckManager interface {
	List(userID string) ([]Deck, error)
	Create(userID string, deck Deck) (string, error)
	Share(userID, deckID string) error
	View(userID, deckID string) (Deck, error)
}
