package main
import (
	"fmt"
	"errors"
)
var firstname string
var lastname string
func main() {
	newname, err := makeName("Jones")
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("Hello, %s\n", newname)
}
func makeName(lname string) (fullname string, err error) {
	err = nil; fullname = ""; firstname = "John"; lastname = lname
	if lastname != "Smith" {
		err = errors.New("Bad last name: " + lname)
	} else {
		fullname = firstname + " " + lastname
	}
	return
}