package main

import (
	"fmt"
	"os"
	"os/exec"
	"log"
	"strconv"
	"io/ioutil"

)

func Clear(){
	//just clears the screen
	cls := exec.Command("clear")
	err := cls.Run()
	if err != nil{
		log.Fatal(err)
	}
}


func Check(er error){
	if er != nil{
		log.Fatal(er)
	}
}


func MakeFile(){
	var filename string
	fmt.Print("filename: ")
	fmt.Scanln(&filename)

	file, err := os.Create(filename)
	Check(err)
	file.Close()
}


func ListDir(){
	//this just list the directory
	dirs, err := ioutil.ReadDir(".")
	Check(err)

	for _, file := range dirs{
		fmt.Println(file.Name())
	}
}

func RemoveDir(){
	var filename string
	fmt.Print("filename: ")
	fmt.Scanln(&filename)
	//deleat the file
	err := os.Remove(filename)
	Check(err)
}


func ReadFile(){
	var filename string
	fmt.Print("filename: ")
	fmt.Scanln(&filename)

	file, err := os.ReadFile(filename)
	Check(err)
	fmt.Println(file)

}
func Menu() int{
	//this will print the menu
	Clear()
	var inp string
	fmt.Println(`
[0] make a file
[1] read a file
[2] delete file
[3] list directory
[4] exit
		`)

	//get a user input
	fmt.Print(": ")
	fmt.Scanln(&inp)

	//temp variable we will return by converting a string to int
	temp, err := strconv.Atoi(inp)
	Check(err)

	return temp
}


func main() {

	for {
		userInput := Menu()
		switch(userInput){
			case 0:
				MakeFile()
			case 1:
				ReadFile()
			case 2:
				RemoveDir()
			case 3:
				ListDir()
			default:
				os.Exit(3)
	}
	}
}