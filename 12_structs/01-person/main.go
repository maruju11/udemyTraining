package main

import "fmt"

type Person struct {
	Name string
	Last string
	Age  int
}

type DoubleAero struct {
	Person
	LicenceToKill bool
}

func (p Person) Greeting() {
	fmt.Println("função-greeting")
}
func (dz DoubleAero) Greeting() {
	fmt.Println("DoubleAero")
}

func main() {
	p1 := Person{
		Name: "James bond",
		Age:  32,
	}
	p2 := DoubleAero{
		Person: Person{
			Name: "Double",
			Age:  1132123,
		},
	}
	p1.Greeting()
	p2.Greeting()
	p2.Person.Greeting()
	fmt.Println(p1.Name, p1.Age)
	fmt.Println(p2.Name, p2.Age)
}
