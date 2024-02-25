package main

import "fmt"
 
 
func main() {
	cards := newDeck()
	fmt.Println(cards.toString())
	fileName := "mycards.txt"
	cards.saveToFile(fileName)
	newCards := newDeckFromFile(fileName)

	fmt.Println(newCards)
}
 