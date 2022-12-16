package utility

import (
	"strings"
)

func Check(err error) {
	if err != nil{
		panic(err)
	}
}

func CalculateGPA(grade map[string]string) int{
	total := 0
	gradeKey := map[string]int{"A": 4, "B": 3, "C":2, "D": 1, "F":0}

	for _, x := range grade{
		gra := gradeKey[x]

		total = total + gra
	}
	return total / len(grade)
}


func CalculateClassAvrage(c Class) float32{
	var avrage float32 = 0

	for _, student := range c.Students{
		cf := CalculateGPA(student.Grades)
		avrage = avrage + float32(cf)
	}
	return avrage / float32(len(c.Students))
}



func ParseGrades(s string) map[string]string{
	finalGrades := make(map[string]string)
	splitSeg := strings.Split(s, ", ")

	counter := 0
	for counter <= len(splitSeg) - 1{
		sepSpaces := strings.Split(splitSeg[counter], " ")

		finalGrades[sepSpaces[0]] = sepSpaces[1]
		counter++
	}

	return finalGrades
}


func ViewGrades(g map[string]string) string{
	master := ""
	for y, x := range g{
		constring := g[y] + " " + g[x] + ", "
		master = master + constring
	}
	return master
}