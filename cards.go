package main // is now an executable type package
import "fmt"

// Auto called from executable
func main1() { // renamed from main so that tests dont try and use this file
	// declare name type = "value"
	var card string = "Ace of Spades" // LONG FORM
	// bool, string, int, float64 <- Basic var types

	// Go compiler will "Figure out" what type of data is assigned to the variable
	card1 := "King of Diamonds" // ABBREVIATED FORMAT

	// Only use := when we are defining the var for the first time, reassigning does not require the :
	card1 = "Five of Diamonds"
	fmt.Println(card)
	fmt.Println(card1)

	card2 := newCard()
	fmt.Println(card2)

	// Slices => []data-type{elements to go in slice}
	cards := []string{newCard(), newCard()}
	fmt.Println(cards)
	// Returns a new slice with the appended value
	cards = append(cards, "Six of Spades")

	//	index, currentElement := itterator
	for i, card := range cards {
		fmt.Println(i, card)
	}

	// GO IS NOT AN OO PROGRAMMING LANGUAGE

}

// declare name() type of data to return {}
func newCard() string {
	// The function initially expects to return no data, we have to tell the compiler that it will return a string
	return "Five of Diamonds x2"
}
