package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

//Cria um novo tipo de "deck"
//que é um slice de string

type deck []string

func (d deck) print(){

    for _, card := range d {
        fmt.Println(card)
    } 
    
}

func newDeck() deck{

    cards := deck{}

    cardSuits := []string{"Spades", "Hearts", "Diamonds", "clubs"}
    cardValues := []string{"Ace ", "Two ", "Three ",  "Four ", "Five ", "Six ", "Seven ", "Eight ", "Nine ", "Ten ", "Valete ", "Dama ", "Rei "}

    for _, suit := range cardSuits{
        for _, value := range cardValues{
            cards = append(cards, value + "of " + suit)
        }
    }

    return cards
}

func (d deck) deal(handsize int) (deck, deck){
    
    hand := d[:handsize]
    d = d[handsize:]

    return hand, d

}


func (d deck) saveDeck(file string){
    
    cardsToSave := []byte(d.toString())

    os.WriteFile(file, cardsToSave, 0666)
}

func (d deck) toString() string{

    return strings.Join([]string(d), ",")
}

func ReadDeck(file string) (deck, error) {

    deckFound, err := os.ReadFile(file)

    if err != nil {
        return nil, err
    }

    slices :=  strings.Split(string(deckFound), ",")
    return deck(slices), nil

}

func (d deck) shuffle() deck{

    source := rand.NewSource(time.Now().UnixNano()) 
    randObj := rand.New(source)

    for i := range d {   
        newPosition := randObj.Intn(len(d)-1)
        d[i], d[newPosition] = d[newPosition], d[i]
    }

    return d

}


