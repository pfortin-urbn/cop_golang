package main

import "fmt"

type Item struct {
	Name string
	Typ int
	Ap int
	Dp int
}

type Player struct {
	Name string
	Level int
	Gold int
	Bag []Item
}

func main() {
	i1 := Item{Name: "gun", Typ: 1, Ap: 20, Dp: 0}
	bag1 := []Item{i1}
	p1 := Player{Name: "John", Level: 1, Gold: 30, Bag: bag1}
	fmt.Println(p1)
}
