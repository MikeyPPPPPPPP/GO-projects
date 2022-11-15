package main


/*
GOAL


1.make an account
	-name
	-password
	-pin
	-ammount

2.view account if you have the password an pin
	-view balence
	-whithdraw
	-deposit

*/


import (
	"fmt"
	"strconv"
	"os"
	"log"
)


//this is a struc for the account
type Account struct {
	username string
	password string
	pin int
	balance float32
}



//this will verify a login and return the index 
func login(acc []Account) int {// it will return the index for the account

	//once we get the username index we will store it here and use it for the rest of the time
	var sec int

	//this is the username check
	fmt.Print("login: ")
	var inp string
	fmt.Scanln(&inp)

	//if all of these are true it will return the index number
	un, pas, pi := false, false, false

	for i := 0; i < len(acc); i++ {


		if inp == acc[i].username {
			un = true
			sec = i
		} else {
			continue
		}
	}
	//if the username dose not exist exit
	if un == false {
		os.Exit(3)
	}

	//check the password
	fmt.Print("Password: ")
	var pass string
	fmt.Scanln(&pass)

	if pass == acc[sec].password {
		pas = true
	} else {
		os.Exit(3)
	}

	//check the pin
	fmt.Print("Pin: ")
	var pinn string
	fmt.Scanln(&pinn)

	//convert the pin to an int
	pinInt, err := strconv.Atoi(pinn)

	if err != nil{
		log.Fatal(err)
	}

	if pinInt == acc[sec].pin {
		pi = true
	} else {
		os.Exit(3)
	}

	if un == true && pas == true && pi == true {
		return sec
	} else {
		os.Exit(3)
	}

	return sec
}

//this will make a struct and add it to the DB
func make_account(acc []Account) []Account{
	//to use DB = make_account(DB)


	//get all the inputs as string
	var info [4]string

	fmt.Println("username, password, pin, ammount")
	for i := 0; i < 4; i++{
		fmt.Print(":")
		//scanner := bufio.NewScanner(os.Stdin)
		//scanner.Scan()

		var scanner string
		fmt.Scanln(&scanner)

		info[i] = scanner
	}

	//we will conver the data accordingly
	p, _ := strconv.Atoi(info[2])
	pin := p


	a, err := strconv.ParseFloat(info[3], 32)
	if err != nil{
		log.Fatal(err)
	}

	amm := float32(a)

	var newUser Account = Account{info[0], info[1], pin, amm}
	
	acc = append(acc, newUser)	

	return acc
}

//print ballance
func show_account(acc []Account, sec int) float32{
	fmt.Println("Current Balence")
	return acc[sec].balance
}


//this will add to the balance
func deposit(acc []Account, sec int) []Account{
	var scanner string
	fmt.Scanln(&scanner)

	am, err := strconv.ParseFloat(scanner, 32)
	if err != nil{
		log.Fatal(err)
	}

	acc[sec].balance = acc[sec].balance + float32(am)
	fmt.Println(acc[sec].balance)

	return acc
}



//this will remove from the balance
func withdraw(acc []Account, sec int) []Account{
	var scanner string
	fmt.Scanln(&scanner)

	am, err := strconv.ParseFloat(scanner, 32)
	if err != nil{
		log.Fatal(err)
	}

	acc[sec].balance = acc[sec].balance - float32(am)


	return acc
}


func main() {


	var DB []Account

	var test Account = Account{"t", "y", 12, 23.0}

	DB = append(DB, test)


	//deposit(DB, 0)


	//fmt.Println(DB[0].balance)



	var sectionNumber int = 44//this is for a logged in user
	fmt.Println(sectionNumber)

	//menu

	for {
		fmt.Println("\n")
		fmt.Println(`
Bank

options:
[0] login
[1] make account
[2] check account
[3] deposit
[4] withdraw
[5] logout

[enter] quit

			`)

		fmt.Print(": ")
		var opt string
		fmt.Scanln(&opt)

		switch opt{
			case "0":
				sectionNumber = login(DB)
				continue
			case "1":
				if sectionNumber != 44 {
					fmt.Println("Log Out first")
					continue
				}
				DB = make_account(DB)
				continue
			case "2":
				if sectionNumber == 44 {
					fmt.Println("Not Logged In")
					continue
				} else {
					am := show_account(DB, sectionNumber)
					fmt.Println(am)
					continue
				}
			case "3":
				if sectionNumber != 44 {
					fmt.Println("Log Out first")
					continue
				}
				deposit(DB, sectionNumber)
				continue

			case "4":
				if sectionNumber != 44 {
					fmt.Println("Log Out first")
					continue
				}
				withdraw(DB, sectionNumber)
				continue


			case "5":
				sectionNumber = 44
				continue
			default:
				os.Exit(3)
		}
	}


	
}

