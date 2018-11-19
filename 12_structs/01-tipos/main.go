package main

import "fmt"

func main() {
	jorge := user{
		name:  "Jorge",
		email: "jorge@gmail.com",
		id:    1324567908,
		adm:   true,
	}
	jj := admin{
		pessoa: user{
			name:  "JJ",
			email: "mail@com",
			id:    11,
			adm:   false,
		},
		nivel: "advanced",
	}
	fmt.Println(jorge.name, jorge.email, jorge.id, jorge.adm)
	fmt.Printf("%T -  -    %p\n", jorge, &jorge)
	fmt.Printf("%s - %s\n", jj.pessoa.name, jj.nivel)
}

type user struct {
	name  string
	email string
	id    int
	adm   bool
}
type admin struct {
	pessoa user
	nivel  string
}
