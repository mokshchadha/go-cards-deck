package main

import "fmt"

//create a new type of deck
//which is a slice of string
type deck []string

func newDeck() deck {
	cards:= deck{}
	cardSuits:= []string {
		"spades" , "diamonds", "hearts", "clubs",
	}
	cardValues := []string {
		"ace", "two", "three", "four",
	}

	for _, suit:= range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+" of "+value)
		}
	}

	return cards
}

func deal(d deck, handSize int) (deck, deck){
	return d[: handSize], d[handSize:]
}

func (d deck) print(){
	for i,card:= range d {
		fmt.Println(i, card)
	}
}