package main

import (
	"fmt"
	c "github.com/skilstak/go/colors"
	//"strings"
	//i "github.com/skilstak/go/input"
	"os"
)

func printPlain() {
	//Prints "Hello World!" to the command line.
	fmt.Println(c.Clear + "Hello World" + c.X)
}
func printRandom() {
	//Prints "Hello World" to the command line in random color.
	fmt.Println(c.Clear + c.Rc() + "Hello World!" + c.X)
}

func printForevah() {
	//Prints "Hello World" in a random color forever in a loop.
	for {
		fmt.Println(c.Rc() + "Hello World" + c.X)
	}
}

func main() {
	name := os.Args[1]
	printMulti("Hello " + getName())
}
func parseArgs() {

}
func getName() string {
	//Gives the var "name" a value and returns it.
	name := "Alexander"
	return name
}

func printMulti(message string) {
	//Prints "Hello World!" To the command line, with each character a random color.
	fmt.Println(c.Clear + c.Multi(message) + c.X)
}
