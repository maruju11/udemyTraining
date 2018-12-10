package main

import "fmt"

// Person comment
type Person struct {
	Name string
	Last string
	Age  int
}

// DoubleAero comment
type DoubleAero struct {
	Person
	LicenceToKill bool
}

// Greeting comment
func (p Person) Greeting() {
	fmt.Println("função-greeting")
}

// Greeting DoubleAero comments
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
