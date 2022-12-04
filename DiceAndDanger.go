/*


make player struct
	attributs map{"":1}

multisided dice sturnt
	side int

dice interface

	roll()

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

type Player struct{
	name string
	attribuits map[string]int
}


type Die struct{
	sides int
}

type Dice interface{
	Roll() int
}

func (d Die) Roll() int{
	randNum := rand.Intn(d.sides)
	return randNum
}




func main(){
	var twenty Dice
	twenty = Die{20}

	var Player1 Player
	Player1 = Player{"Michael", map[string]int{"Strength":0, "Constitution":0, "Dexterity":0, "Intelligence":0, "Wisdom":0, "Charisma":0}}

	for x, _  := range Player1.attribuits{
		Player1.attribuits[x] = twenty.Roll()
	}

	for y, x := range Player1.attribuits{
		fmt.Print(y+ ": ")
		fmt.Println(x)
	}
}







