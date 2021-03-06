package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
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

func (d deck) saveToFile(fname string) {
	deckStringSlice := []string(d)
	deckString := strings.Join(deckStringSlice, "\n")
	deckByte := []byte(deckString)
	ioutil.WriteFile(fname, deckByte, 0666)
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

func (d deck) shuffle() {
	unix := time.Now().UnixNano()
	source := rand.NewSource(unix)
	newrand := rand.New(source)
	for i := range d {
		newPosition := newrand.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
