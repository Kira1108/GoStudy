package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	colors := []string{"Heart", "Spade", "Diamond", "Club"}
	numbers := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	d := deck{}

	for _, color := range colors {
		for _, num := range numbers {
			d = append(d, color+num)
		}
	}
	d = append(d, "Joker black")
	d = append(d, "Joker red")
	return d
}

func deal(d deck, n int) (deck, deck) {
	return d[:n], d[n:]
}

func (d deck) saveToFile() {
	deckStringSlice := []string(d)
	deckString := strings.Join(deckStringSlice, "\n")
	deckByte := []byte(deckString)
	ioutil.WriteFile("cards.txt", deckByte, 0666)
}

func readDeckFromFile(fname string) deck {
	deckByte, err := ioutil.ReadFile(fname)
	if err != nil {
		fmt.Println("Error read deck from file")
		return nil
	}
	deckString := string(deckByte)
	deckStringSlice := strings.Split(deckString, "\n")
	thisDeck := deck(deckStringSlice)
	fmt.Println("Successfully reading deck from file")
	return thisDeck
}
