package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

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
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// func deal(d deck, handSize int) (deck, deck){
// 	return d[: handSize], d[handSize:]
// }

func (d deck) print(){
	for i,card:= range d {
		fmt.Println(i, card)
	}
}

func (d deck)toString() string {
	str:=strings.Join([]string(d), ",")
	return str
}

func (d deck)saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error: while reading ", fileName)
		fmt.Println(err)
		os.Exit(1)
	}

	str:=string(bs)
	ss:= strings.Split(str, ",")
	return deck(ss)
}

func (d deck) shuffle(){
	source:=   rand.NewSource(time.Now().UnixMicro())
	r:= rand.New(source)
	for i := range d{
		newPosition := r.Intn(len(d) -1)
		d[i] , d[newPosition] = d[newPosition], d[i]
	}
}