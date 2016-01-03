package main

import (
	"fmt"
	"time"
	"math/rand"
)

type Item struct {
	Name string
	Typ int
	Ap int
	Dp int
}

func (i *Item) Shoot() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(i.Ap)
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
	fmt.Println("He shoots and hits for:", i1.Shoot())
}
