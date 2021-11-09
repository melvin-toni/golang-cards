package main

func main() {
	/*
		how to run go in console
		- go run main.go deck.go
		- go build main.go ----then----> main
	*/

	/*
		2 ways to initialize data :
		- var card string = "Ace of Spades"
		- card := "Ace of Spades"
	*/

	// Assign value
	// card := newCard()
	// fmt.Println(card)

	// Loop with slice
	// cards := []string{"Ace of Diamonds", newCard()}
	// cards = append(cards, "Six of Spades")
	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }

	// Deck data type and receiver
	// cards := newDeck()
	// cards.print()

	// Assign value to multiple variables
	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// remainingCards.print()

	/*
		It will saved into local storage using ioutil pkg and WriteFile func
		to do that we convert slice of string into byte slice, because second argument of
		WriteFile func is []byte, here is the diagram
		deck -> []string -> string -> []byte
	*/
	// cards.saveToFile("my_cards")

	/*
		It will load file from local storage and print it
	*/
	// cards := newDeckFromFile("my_cards")
	// cards.print()

	/*
		It will shuffle the card
	*/
	cards := newDeck()
	cards.shuffle()
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
