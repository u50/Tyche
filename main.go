package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Deck struct {
	Cards []Card
}

type Card struct {
	Suit string
	Rank string
}

func shuffleDeck() Deck {
	deck := Deck{
		[]Card{
			{"Heart", "Two"},
			{"Heart", "Three"},
			{"Heart", "Four"},
			{"Heart", "Five"},
			{"Heart", "Six"},
			{"Heart", "Seven"},
			{"Heart", "Eight"},
			{"Heart", "Nine"},
			{"Heart", "Ten"},
			{"Heart", "Jack"},
			{"Heart", "Queen"},
			{"Heart", "King"},
			{"Heart", "Ace"},
			{"Diamond", "Two"},
			{"Diamond", "Three"},
			{"Diamond", "Four"},
			{"Diamond", "Five"},
			{"Diamond", "Six"},
			{"Diamond", "Seven"},
			{"Diamond", "Eight"},
			{"Diamond", "Nine"},
			{"Diamond", "Ten"},
			{"Diamond", "Jack"},
			{"Diamond", "Queen"},
			{"Diamond", "King"},
			{"Diamond", "Ace"},
			{"Club", "Two"},
			{"Club", "Three"},
			{"Club", "Four"},
			{"Club", "Five"},
			{"Club", "Six"},
			{"Club", "Seven"},
			{"Club", "Eight"},
			{"Club", "Nine"},
			{"Club", "Ten"},
			{"Club", "Jack"},
			{"Club", "Queen"},
			{"Club", "King"},
			{"Club", "Ace"},
			{"Spade", "Two"},
			{"Spade", "Three"},
			{"Spade", "Four"},
			{"Spade", "Five"},
			{"Spade", "Six"},
			{"Spade", "Seven"},
			{"Spade", "Eight"},
			{"Spade", "Nine"},
			{"Spade", "Ten"},
			{"Spade", "Jack"},
			{"Spade", "Queen"},
			{"Spade", "King"},
			{"Spade", "Ace"},
		},
	}
	rand.Seed(time.Now().Unix())

	rand.Shuffle(len(deck.Cards), func(i, j int) {
		deck.Cards[i], deck.Cards[j] = deck.Cards[j], deck.Cards[i]
	})

	return deck
}

func (card *Card) getSuit() string {
	return card.Suit
}

func (card *Card) getRank() string {
	return card.Rank
}

func main() {
	deck := shuffleDeck()
	fmt.Println(deck.Cards[0].getSuit())
	fmt.Println(deck.Cards[0].getRank())
}
