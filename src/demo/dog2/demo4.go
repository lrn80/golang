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
	// 示例1。
	var dog1 *Dog
	if dog1 == nil {
		fmt.Println("The first dog is nil.") // 输出 The first dog is nil.
	} else {
		fmt.Println("The first dog is not nil.")
	}

	dog2 := dog1
	if dog2 == nil {
		fmt.Println("The second dog is nil.") // 输出 The second dog is nil.
	} else {
		fmt.Println("The second dog is not nil.")
	}
	var pet Pet = dog2

	if pet == nil {
		fmt.Println("The pet is nil.")
	} else {
		fmt.Println("The pet is not nil.") // 输出 The pet is not nil.
	}
}
