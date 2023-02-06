package main

import (
	"fmt"
)

var pl = fmt.Println

type Animal interface {
	AngrySound()
	HappySound()
}

type Cat string

func (c Cat) Name() string {
	return string(c)

}

func (c Cat) Attack() {
	pl("Cat attack its prey")

}

func (c Cat) AngrySound() {
	pl("cat says hiss")

}

func (c Cat) HappySound() {
	pl("cat says Purr")

}

func main() {

	var kitty Animal
	kitty = Cat("Kitty")
	kitty.AngrySound()

	var kitty2 Cat = kitty.(Cat)

	kitty2.Attack()
	pl("v=cats name :", kitty2.Name())

}
