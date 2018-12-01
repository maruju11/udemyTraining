package main

import "fmt"

func main() {
	//slices usando make: reference type > aponta para um endereço com data
	student := make([]string, 20, 35)
	students := make([]string, 20, 35)
	student[0] = "00Jorge"
	student[19] = "19Jorge"
	student = append(student, "Jorge20")
	fmt.Println(student)
	fmt.Println(students)
	fmt.Println(student == nil)

	// slices usando var o header está apontando para nada. não existe endereço nem data associado a ela
	var s1tudent []string
	var s1tudents [][]string
	//s1tudent[0] = "Karla" - // nao roda, pois dá index out of range
	s1tudent = append(s1tudent, "1Jorge")
	fmt.Println(s1tudent)
	fmt.Println(s1tudents)
	fmt.Println(s1tudent == nil)

	//a mesma coisa aqui e o array nao está pronto
	s2tudent := []string{}
	s2tudents := [][]string{}
	//s2tudent[0] = "Karla" - nao roda, pois dá index out of range
	s2tudent = append(s2tudent, "2Jorge")
	fmt.Println(s2tudent)
	fmt.Println(s2tudents)
	fmt.Println(s2tudent == nil)

}
