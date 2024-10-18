package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	cards := []string{newCard(),"Ace of Diamonds"}
	fmt.Println(cards)
}

func newCard() string {
	return "Five of Diamonds"
}