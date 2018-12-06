package main

import "fmt"

type user struct {
	name  string
	email string
}

func (u user) notify() {
	fmt.Printf("Foi enviado um e-mail para: %s<%s>\n", u.name, u.email)
}
func (u *user) changeMail(em string) {
	u.email = em
	fmt.Printf("e-mail alterado: %s <%s>\n", u.name, u.email)
}
func main() {
	u1 := user{"Jorge", "mail@gmail.com"}
	u1.notify()
	u2 := &user{"Lisa", "lisa@gmail.com"}
	u2.notify()
	u2.changeMail("alterado@gmail.com")
}
