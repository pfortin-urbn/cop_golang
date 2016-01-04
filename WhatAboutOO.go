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

func main() {
	i1 := Item{Name: "gun", Typ: 1, Ap: 20, Dp: 0}
	fmt.Println("He shoots and hits for", i1.Shoot(), "Hit Points.")
}
