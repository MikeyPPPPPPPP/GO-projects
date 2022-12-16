package main

import (
	"fmt"
	"os"
	"bufio"
	us "example.com/packages/utility"
)


//hash this latter
func Login() bool{
	var password string
	fmt.Print("Password: ")
	fmt.Scanln(&password)
	if password == "password"{
		return true
	}
	return false
}

func AddClass(s us.School) us.School{
	var newClass us.Class

	var teacherName string
	fmt.Print("Teacher Name:")
	fmt.Scanln(&teacherName)

	newClass.Teacher = teacherName

	for {
		if us.AddStudentYN(){
			var studentFN string
			var studentLN string
			var grades string


			fmt.Print("Student first name: ")
			fmt.Scanln(&studentFN)
			fmt.Print("Student last name: ")
			fmt.Scanln(&studentLN)
			fmt.Print("Student grades\nex: math A, Sience B\n:")
			inp := bufio.NewScanner(os.Stdin)
			inp.Scan()
			grades = inp.Text()

			newClass.Students = append(newClass.Students, us.Student{studentFN, studentLN, us.ParseGrades(grades), us.CalculateGPA(us.ParseGrades(grades))})
			fmt.Println("Student Added")
		} else {
			break
		}
	}

	s.Classes = append(s.Classes, newClass)
	fmt.Println("Class added")
	return s
}



func ViewClass(c us.Class) {

	inp := us.ViewClassMenu()
	switch inp{
		case 1://view students
			fmt.Println("viewing students")
			for _, student := range c.Students{
				fmt.Println(student.Name + " " + student.LastName + " " + us.ViewGrades(student.Grades))
			}
		case 2://student
			var stuname string
			fmt.Print(":")
			fmt.Scanln(&stuname)

			for _, student := range c.Students{
				if student.Name == stuname{
					fmt.Println(student.Name + " " + student.LastName + " " + us.ViewGrades(student.Grades))
				}
			}
		case 3://class avrage
			fmt.Printf("Class Avrage is %g", us.CalculateClassAvrage(c))

	}
}

func PickClass(s us.School) us.Class{
	var teacher string
	fmt.Print("enter Teacher Name: ")
	fmt.Scanln(&teacher)

	default_class :=  s.Classes[0]
	for _, class := range s.Classes{
		if teacher == class.Teacher{
			return class
		}
	}
	fmt.Println("Class dose not exist so here is the first one!")
	return default_class
	

}

func main(){
	//var first us.Class
	//first.Teacher = "Michael"
	//first.Students = append(first.Students, us.Student{"sash", "provenzano", map[string]string{"math":"A", "science":"B"}, 4})

	var ms us.School

	//temp
	var first us.Class
	first.Teacher = "Michael"
	first.Students = append(first.Students, us.Student{"sash", "prov", map[string]string{"math":"A", "science":"B", "history":"B"}, 4})

	ms.Classes = append(ms.Classes, first)

	if Login(){
		for {
			o := us.MainMenu()
			if o == 1{
				ms = AddClass(ms)
			} else if (o == 2){
				ViewClass(PickClass(ms))

			} else {
				break
			}
		}


	}

}