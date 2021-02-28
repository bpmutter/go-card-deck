package main
import (
	"fmt"
)

func main() {
	cards := newDeck()

	// hand, remaining := deal(cards, 5)
	// hand.print()
	// remaining.print()
	cards.saveToFile("my_cards.txt")
	deck := newDeckFromFile("my__cards.txt")

		fmt.Println(deck.stringify())
		fmt.Println("success!")

}

