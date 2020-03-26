package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

// create a new type of 'deck'
// which is a slice of strings
type deck []string

// define a method of deck types
// func (var_name type)
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	// using _ instead of a var name (like i) it tells Go that we dont really want to use the values that will be stored there
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

//func (d deck) shuffleDeck() deck {
//
//}

func deal(d deck, handSize int) (deck, deck) {
	// slice[startIndexIncluding : upToNotIncluding]
	// slice[0:2]
	// fruits := []string{"apple", "banana", "grape", "orange"}
	// fmt.Println(fruits[:2])
	// [apple banana]
	// The hand, Whats left in the deck
	return d[:handSize], d[handSize:]
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
	// converting types -> []type(whatToConvert)
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func newDeckFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		// Option #1 print out error and call newDeck()
		fmt.Println("Failed to read file: ", err)
		return newDeck()
		// Option #2 log the error  and quit the program
		// fmt.Println("Failed to read the file: ", err)
		// os.Exit(2)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	for i := range d {
		source := rand.NewSource(time.Now().UnixNano())
		r := rand.New(source)
		index := r.Intn(len(d) - 1)
		d[i] = d[index]
		d[index] = d[i]
		// or d[i], d[index] = d[index], d[i]
	}
}
