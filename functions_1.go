package main
import (
	"fmt"
)
var firstname string
var lastname string
func main() {
	newname, err := makeName("Smith")
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("Hello, %s\n", newname)
}
func makeName(lname string) (string, error) {
	firstname = "John"
	lastname = lname
	if lastname != "Smith" {
		return "", fmt.Errorf("Bad last name: %s\n", lastname)
	}
	return firstname + " " + lastname, nil
}
