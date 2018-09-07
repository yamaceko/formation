package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

//create a new type equivalat a class in c#
//it's a slice of strings

type deck []string

func newDeck() deck {

	cards := deck{}
	cardSuits := []string{"Pick", "Careaux", "Trefles", "Coeurs"}
	cardValues := []string{"As", "Un", "Deux", "Trois", "Quattre"}

	for _, suits := range cardSuits {

		for _, values := range cardValues {
			cards = append(cards, values+" de "+suits)
		}
	}
	return cards

}

func (d deck) print() {

	for i, d := range d {
		fmt.Println(i, d)
	}
}

func deal(d deck, handSize int) (deck, deck) {

	//mean ex: for d[10] and handsize=3
	//1: we've 0 to 2
	//2: we've 3 to 9

	return d[:handSize], d[handSize:]

}

//convert deck to string sequence values
func (d deck) toString() string {

	return strings.Join([]string(d), ",")
}

//save deck into file
func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}


//create deck from file saved
func newDeckFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error :", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}


//it's use to shuffle a deck
func (d deck) shuffle() {

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]

	}
}
