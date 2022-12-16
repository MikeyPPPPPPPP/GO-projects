package utility

import (
	"fmt"
	"strconv"
)

func MainMenu() int{
	fmt.Println(`
1. add class
2. view class
`)	
	var option string
	fmt.Print(":")
	fmt.Scanln(&option)
	con, err := strconv.Atoi(option)
	Check(err)
	return con
}


func AddStudentYN() bool{
	var question string
	fmt.Print("add another student (y/n):")
	fmt.Scanln(&question)
	if question == "y"{
		return true
	}
	return false
}

func ViewClassMenu() int {
	fmt.Println(`
1. view students
2. view student
3. viev class avrage
`)	
	var option string
	fmt.Print(":")
	fmt.Scanln(&option)
	con, err := strconv.Atoi(option)
	Check(err)
	return con
}
