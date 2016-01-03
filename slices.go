package main

import "fmt"

func main() {
	elements := []string{"cat", "dog", "bird"}
	elements = append(elements, "fish", "snake")
	fmt.Println(elements)
	fmt.Println(elements[3])
	fmt.Println(len(elements))

	//Detete an element [3]
	itemToDelete := 3
	result := append(elements[:itemToDelete], elements[itemToDelete+1:]...)
	fmt.Println(result)

	inserted := []string{"fish2"}
	result2 := append(result[:3], append(inserted, result[3:]...)...)
	fmt.Println(result2)
}
