package struct2

import "fmt"

type Person struct {
	name string
	age int
	skinColor string

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
	blackPerson := Person{name: "xiaohei", age: 20, skinColor: "black"}
	fmt.Println(blackPerson) // 输出 {xiaohei 20 black}

	blackPerson.set("white")
	fmt.Println(blackPerson) // 输出 {xiaohei 20 black}

	blackPerson.setColor("red")
	fmt.Println(blackPerson) // {xiaohei 20 red}
}
