package main

import "fmt"

type Person struct {
	name string
	age int
	skinColor string
	family Family
}

type Family struct {
	familyName string
}

func (ac Family) print()  {
	fmt.Printf("print familyName: %s", ac.familyName)
}

func (ac Family) String() string {
 	return fmt.Sprintf("familyName: %s",ac.familyName)
}

func (ac Person) String() string {
	return fmt.Sprintf("name: %s age: %s skinColor: %s family: %s", ac.name, ac.age, ac.skinColor, ac.family)
}

func (ac Person) printfSkinColor() {
	fmt.Printf("skinColor%s", ac.skinColor)
}

func (ac Person) set(skinColor string)  {
	ac.skinColor = skinColor
}

func (ac *Person) setColor(skinColor string)  {
	ac.skinColor = skinColor
}

func main() {
	Corleone := Family{familyName: "Corleone"}

	person := Person{name: "Tom", age: 40, skinColor: "black", family: Corleone}

	fmt.Printf("%s", person) // 输出 name: Tom age: %!s(int=40) skinColor: black family: familyName: Corleone

	//blackPerson := Person{name: "xiaohei", age: 20, skinColor: "black"}
	//fmt.Println(blackPerson) // 输出 {xiaohei 20 black}
	//
	//blackPerson.set("white")
	//fmt.Println(blackPerson) // 输出 {xiaohei 20 black}
	//
	//blackPerson.setColor("red")
	//fmt.Println(blackPerson) // {xiaohei 20 red}
}
