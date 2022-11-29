package main

/*
BASIC GOALS

1.make a deck of cards
2.shuffle deck of cards
3.print deck of cards


*/




import (
	"fmt"
	"strconv"
	"math/rand"
	"time"
)

type Card struct{
	suit string
	num string
}

type Deck struct{
	card []Card
}

//this will populat a deck with an array of cards
func makeNewDeck(deck Deck) Deck{
	suits := [4]string{"D","H","C","S"}
	for _, x := range suits{
		for i := 1; i < 14; i++{
			if i == 1{
				deck.card = append(deck.card, Card{x, "A"})
			} else if i == 11{
				deck.card = append(deck.card, Card{x, "J"})
			} else if i == 12{
				deck.card = append(deck.card, Card{x, "Q"})
			} else if i == 13{
				deck.card = append(deck.card, Card{x, "K"})
			} else {
				deck.card = append(deck.card, Card{x, strconv.Itoa(i)})
			}
		}		
	}
	return deck

}


//shuffles the deck in a random order
func shuffleCards(deck Deck) Deck{
	var temp []Card

	rand.Seed(time.Now().UnixNano())
	for {
		if len(deck.card) == len(temp){
			break
		}

		randomCard := deck.card[rand.Intn(52)]

		var inside bool
		inside = false

		for _, i := range temp{
			if randomCard == i{
				inside = true
			}
		}
		if inside == false{
			temp = append(temp, randomCard)
		}
	}

	deck.card = temp
	return deck
}

//shows card in a nicer way 
func printDeck(deck Deck){
	for _, i := range deck.card{
		fmt.Print(i.suit + i.num + " ")
	}
}

func main(){
	var numDeck Deck
	fmt.Println("Making deck")
	numDeck = makeNewDeck(numDeck)
	fmt.Println("shuffling deck")
	//fmt.Println(numDeck)
	numDeck = shuffleCards(numDeck)
	fmt.Println("printing deck")
	printDeck(numDeck)


}	