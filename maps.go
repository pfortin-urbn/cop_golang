package main

import "fmt"

var m map[string]int

func main() {
	commits := map[string]int{
		"rsc": 3711,
		"r":   2138,
		"gri": 1908,
		"adg": 912,
	}

	fmt.Println(commits)
}
