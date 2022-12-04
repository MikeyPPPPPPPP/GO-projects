/*
player
	health
	inventor
	attack
	position

monster
	health
	attack

map 


l r

*/



package main


import (
	"fmt"
	"time"
	"math/rand"
	"os"
)


func init(){
	rand.Seed(time.Now().UnixNano())
}


type Player struct{
	health int
	inventor []string
	attack int
	position int
}

type Monster struct{
	health int
	attack int
}


func Battle(p Player, m Monster) Player{
	if p.attack > m.health {
		p.attack++
		return p
	}
	p.health = p.health - m.attack
	return p

}
func MainMenu() bool{
	var inp string
	fmt.Print("play y/n")
	fmt.Scanln(&inp)
	if inp == "y"{
		return true
	}
	return false
}

func main(){
	s := MainMenu()
	if s != true{
		os.Exit(3)
	}

	var play Player
	play = Player{4, []string{"gun"}, 4, 1}
	for {
		var ag string
		fmt.Print("play agian? ")
		fmt.Scanln(&ag)
		if ag == "n"{
			os.Exit(3)
		}
		var mon Monster
		mon = Monster{rand.Intn(play.health+7), rand.Intn(play.attack+5)}

		fmt.Println("you might need to fight")
		
		if rand.Intn(5) == 1{
			play = Battle(play, mon)
			if play.health < 1{
				fmt.Println("you were killed fish")
				os.Exit(3)
			} else {
				fmt.Println("you won")
			}


		} else {
			fmt.Println(play.attack)
		}
	}
}



