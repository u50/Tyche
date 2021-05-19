package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Deck struct {
	Cards []Card
}

type Hand struct {
	First  Card
	Second Card
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

func (deck *Deck) drawCard() (Card, Deck) {
	card := deck.Cards[0]
	newDeck := append(deck.Cards[:0], deck.Cards[1:]...)

	return card, Deck{newDeck}
}

func (deck *Deck) drawHand() (Hand, Deck) {
	first, _ := deck.drawCard()
	second, newDeck := deck.drawCard()

	return Hand{first, second}, newDeck
}

func (hand *Hand) getHand() []Card {
	return []Card{hand.First, hand.Second}
}

func (deck *Deck) getTheFlop() ([]Card, Deck) {
	first, _ := deck.drawCard()
	second, _ := deck.drawCard()
	third, newDeck := deck.drawCard()

	return []Card{first, second, third}, newDeck
}

func (deck *Deck) getTheTurn() (Card, Deck) {
	fourth, newDeck := deck.drawCard()

	return fourth, newDeck
}

func (deck *Deck) getTheRiver() (Card, Deck) {
	fifth, newDeck := deck.drawCard()

	return fifth, newDeck
}

func calculateScore(hand Hand, communityCards []Card) int {
	return 0
}

func playRound() {
	// shuffle
	deck := shuffleDeck()

	// deal hands
	playerOneHand, deck := deck.drawHand()
	playerTwoHand, deck := deck.drawHand()

	// show community cards
	communityCards := []Card{}
	communityCards, deck = deck.getTheFlop()

	// show fourth
	fourth, deck := deck.getTheTurn()
	communityCards = append(communityCards, fourth)

	// show fifth
	fifth, deck := deck.getTheRiver()
	communityCards = append(communityCards, fifth)

	// calculate winner
	playerOneScore := calculateScore(playerOneHand, communityCards)
	playerTwoScore := calculateScore(playerTwoHand, communityCards)

	if playerOneScore > playerTwoScore {
		fmt.Println("Player One Won")
	}
	if playerTwoScore > playerOneScore {
		fmt.Println("Player Two Won")
	} else {
		fmt.Println("Draw!")
	}
}

func main() {
	playRound()
}
