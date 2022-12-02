/*

making Fisher and Yates algorithm

*/



package main

import (
	"fmt"
	"time"
	"math/rand"
)

func init(){
	rand.Seed(time.Now().UnixNano())
}

func Swap(obj []string, x int, y int) []string{
	temp := obj[x]
	obj[x] = obj[y]
	obj[y] = temp
	return obj
}

func main(){
	var master = []string{"m", "i","c","h","a","e","l"}
	var tempSlice []string
	var finalSlice []string
	tempSlice = master

	var counter int = len(master)

	for x := 0; x< len(master); x++{
		//pick a random number based on the slice length
		randomNumber := rand.Intn(counter)

		//append thea index to the final slice 
		finalSlice = append(finalSlice, tempSlice[randomNumber])

		//swap randomly picked element with the last element
		tempSlice = Swap(tempSlice, randomNumber, counter - 1)

		counter--
	}

	fmt.Println(finalSlice)


}