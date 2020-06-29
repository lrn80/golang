package main

import "fmt"

type Pet interface {
	Name() string
	Category() string
}

type Dog struct {
	name string // 名字。
}

func (dog *Dog) SetName(name string) {
	dog.name = name
}

func (dog Dog) Name() string {
	return dog.name
}

func (dog Dog) Category() string {
	return "dog"
}

func main() {

	dog := Dog{"little pig"}
	var pet Pet = dog
	fmt.Println(pet == dog)
	dog.SetName("monster")

	fmt.Printf("%T %v\n", pet, pet) // 输出: main.Dog {little pig}

	fmt.Printf("%T %v", dog, dog) // 输出: main.Dog {monster}


}
