package main

import (
	"fmt"
	c "github.com/skilstak/go/colors"
	//i "github.com/skilstak/go/input"
)

func printRandom() {
	//Prints "Hello World" to the command line in random color.
	fmt.Println(c.Clear + c.Rc() + "Hello World!" + c.X)
}

func main() {
	//Prints "Hello World" in a random color forever in a loop.
	for {
		fmt.Println(c.Rc() + "Hello World" + c.X)
	}
}

func main1() {
	//Prints "Hello World!" to the command line.
	fmt.Println(c.Clear + "Hello World" + c.X)
}

func printMulti() {
	//Prints "Hello World!" To the command line, with each character a random color.
	fmt.Println(c.Clear + c.Multi("Hello World") + c.X)
}
