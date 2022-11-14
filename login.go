


//3 time login

package main


import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"log"
)



//takes in a list 
func menu(lis []int) {
	for {
		fmt.Println(`
[1] print list
[2] add to list
[3] exit
			`)

		fmt.Print(": ")
		inp := bufio.NewScanner(os.Stdin)
		inp.Scan()

		switch inp.Text() {
			case "1":
				fmt.Println(lis)
			case "2":
				fmt.Print(": ")
				adder := bufio.NewScanner(os.Stdin)
				adder.Scan()

				//string to in conversion
				intVar, err := strconv.Atoi(adder.Text())

				//handle the err
				if err != nil{
					log.Fatal(err)
				}

				//just append it to the string
				lis = append(lis, intVar)

			default:
				os.Exit(3)
		}
	}
}

func main(){
	//this is the password
	const password string = "Password"

	//login attempts
	var attempts int = 0

	//the list for the menu
	var list []int

	for {
		//check if attempt is not 3
		if attempts == 3 {
			//exit the program
			os.Exit(3)
		} else {
			//get a user input 
			fmt.Print("password: ")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()

			//if the input is the same as the password break
			if scanner.Text() == password {
				break
			} else {

				//if the password is wrong add 1 to the attempt
				attempts++
			}

		}
	}

	menu(list)



}