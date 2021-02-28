package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"os"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	SUITS := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	VALUES := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	
	for _, suit := range SUITS{
		for _, value := range VALUES{
			cards = append(cards, value + " of " + suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
	fmt.Println("---")
}

func deal(cards deck, numCards int) (deck, deck){

		topOfDeck := cards[0:numCards]
		restOfDeck := cards[numCards:]
		return topOfDeck, restOfDeck
	
}

func (d deck) stringify() string{
	
	deckString := strings.Join([]string(d), ", ")
	return deckString
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.stringify()), 0666)
}

func newDeckFromFile(fileName string) deck {
	byteCards, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	cards := strings.Split(string(byteCards), ", ")
	newDeck := deck(cards)
	return newDeck
}