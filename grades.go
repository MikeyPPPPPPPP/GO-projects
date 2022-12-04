package main





/*


make a student

student
	grades
		a
		b
		c
		d

	name
	attendence 



struct student 

interface student
	calculate gpa

*/

import (
	"fmt"
)
type ReportCard interface{
	calculateGPA() int
}

type Student struct{
	name string
	attendence int
	grades [5]string
}




func (s Student) calculateGPA() int{
	total := 0
	gradeKey := map[string]int{"A": 4, "B": 3, "C":2, "D": 1, "F":0}

	for _, x := range s.grades{
		total = total + gradeKey[x]
	}
	return total / 5
}

func main(){

	var rp ReportCard

	rp = Student{"Michael", 100, [5]string{"A", "B", "C", "D", "F"}}
	fmt.Println(rp.calculateGPA())

}