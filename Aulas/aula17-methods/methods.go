package main

import (
	"fmt"
)

type user struct {
	name string
	age  uint8
}

func (u user) identity() {
	fmt.Printf("Identyfing User %s at Data Base\n", u.name)
}

func (u user) identifyAge() uint8 {
	fmt.Println("His/Her age is", u.age)
	return u.age
}
func (u *user) agePlus() {
	u.age++
}

func main() {
	fmt.Println("Methods")

	user1 := user{"Kauã", 17}
	user1.identity()
	user1.identifyAge()

	fmt.Println("------------------------------")

	user2 := user{"Pedro", 20}
	user2.identity()
	user2.identifyAge()

	user2.agePlus()
	fmt.Println(user2.age)
}
