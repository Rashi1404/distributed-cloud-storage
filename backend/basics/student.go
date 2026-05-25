package main

import "fmt"

type student struct {
	Name   string
	Branch string
	CGPA   float64
}

func main() {
	student1 := student{
		Name:   "Rashi",
		Branch: "CSE",
		CGPA:   8.6,
	}
	student2 := student{
		Name:   "Kashish",
		Branch: "IT",
		CGPA:   8.2,
	}
	fmt.Println(student1.Name)
	fmt.Println(student1.Branch)
	fmt.Println(student1.CGPA)

	fmt.Println("------------------------")
	fmt.Println(student2.Name)
	fmt.Println(student2.Branch)
	fmt.Println(student2.CGPA)
}
