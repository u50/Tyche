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

type Suit struct {
	Hearts   string
	Diamonds string
	Clubs    string
	Spades   string
}

type Rank struct {
	Two   string
	Three string
	Four  string
	Five  string
	Six   string
	Seven string
	Eight string
	Nine  string
	Ten   string
	Jack  string
	Queen string
	King  string
	Ace   string
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

type Score struct {
	RoyalFlush    bool
	StraightFlush bool
	FourOfAKind   bool
	FullHouse     bool
	Flush         bool
	Straight      bool
	ThreeOfAKind  bool
	TwoPair       bool
	Pair          bool
	HighCard      Card
}

func (score *Score) highCard(hand Hand, communityCards []Card) Card {
	cardRanks := []string{}
	for _, card := range communityCards {
		// Ace is best, early return :)
		if card.getRank() == "Ace" {
			return card
		}
		cardRanks = append(cardRanks, card.getRank())
	}

	return Card{"Hearts", "Ace"}
}

func (score *Score) pair(hand Hand, communityCards []Card) bool {
	cardRanks := []string{}
	for _, card := range communityCards {
		cardRanks = append(cardRanks, card.getRank())
	}

	return true
}

func (score *Score) twoPairs(hand Hand, communityCards []Card) bool {
	cardRanks := []string{}
	for _, card := range communityCards {
		cardRanks = append(cardRanks, card.getRank())
	}

	return true
}

func (score *Score) threeOfAKind(hand Hand, communityCards []Card) bool {
	cardRanks := []string{}
	for _, card := range communityCards {
		cardRanks = append(cardRanks, card.getRank())
	}

	return true
}

func (score *Score) straight(hand Hand, communityCards []Card) bool {
	cardRanks := []string{}
	for _, card := range communityCards {
		cardRanks = append(cardRanks, card.getRank())
	}

	return true
}

func (score *Score) flush(hand Hand, communityCards []Card) bool {
	cardRanks := []string{}
	for _, card := range communityCards {
		cardRanks = append(cardRanks, card.getRank())
	}

	return true
}

func (score *Score) fullHouse(hand Hand, communityCards []Card) bool {
	cardRanks := []string{}
	for _, card := range communityCards {
		cardRanks = append(cardRanks, card.getRank())
	}

	return true
}

func (score *Score) fourOfAKind(hand Hand, communityCards []Card) bool {
	cardRanks := []string{}
	for _, card := range communityCards {
		cardRanks = append(cardRanks, card.getRank())
	}

	return true
}

func (score *Score) straightFlush(hand Hand, communityCards []Card) bool {
	cardRanks := []string{}
	for _, card := range communityCards {
		cardRanks = append(cardRanks, card.getRank())
	}

	return true
}

func (score *Score) royalFlush(hand Hand, communityCards []Card) bool {
	cardRanks := []string{}
	for _, card := range communityCards {
		cardRanks = append(cardRanks, card.getRank())
	}

	return true
}

func calculateScore(hand Hand, communityCards []Card) (int, Card) {
	score := Score{}

	// Royal Flush?
	if score.royalFlush(hand, communityCards) {
		return 10, score.highCard(hand, communityCards)
	}

	// Straight Flush?
	if score.straightFlush(hand, communityCards) {
		return 9, score.highCard(hand, communityCards)
	}

	// Four of a Kind?
	if score.fourOfAKind(hand, communityCards) {
		return 8, score.highCard(hand, communityCards)
	}

	// Full House?
	if score.fullHouse(hand, communityCards) {
		return 7, score.highCard(hand, communityCards)
	}

	// Flush?
	if score.flush(hand, communityCards) {
		return 6, score.highCard(hand, communityCards)
	}

	// Straight?
	if score.straight(hand, communityCards) {
		return 5, score.highCard(hand, communityCards)
	}

	// Three of a Kind?
	if score.threeOfAKind(hand, communityCards) {
		return 4, score.highCard(hand, communityCards)
	}

	// Two Pair?
	if score.twoPairs(hand, communityCards) {
		return 3, score.highCard(hand, communityCards)
	}

	// Pair?
	if score.pair(hand, communityCards) {
		return 2, score.highCard(hand, communityCards)
	}

	// High Card?
	return 1, score.highCard(hand, communityCards)
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
	playerOneScore, playerOneHighCard := calculateScore(playerOneHand, communityCards)
	playerTwoScore, playerTwoHighCard := calculateScore(playerTwoHand, communityCards)

	if playerOneScore > playerTwoScore {
		fmt.Println("Player One Won")
	}
	if playerTwoScore > playerOneScore {
		fmt.Println("Player Two Won")
	} else {
		if playerOneHighCard > playerTwoHighCard {
			fmt.Println("Player One Won")
		} else {
			fmt.Println("Player Two Won")
		}
	}
}

func main() {
	playRound()
}
