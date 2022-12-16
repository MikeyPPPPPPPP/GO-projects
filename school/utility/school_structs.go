package utility



type School struct{
	Classes []Class
}

type Student struct{
	Name string
	LastName string

	Grades map[string]string //[class] grade
	GPA int
}

type Class struct{
	Teacher string

	Students []Student
}