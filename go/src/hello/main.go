package main

import (
	"fmt"
	c "github.com/skilstak/go/colors"
	//i "github.com/skilstak/go/input"
)

func printRandom() {
	//Prints "Hello World" to the command line in random color.
	fmt.Println(c.Clear + c.Rc() + "Ryan wuz hear" + c.X)
}

func main() {
	//Prints "Hello World" in a random color forever in a loop.
	for {
		fmt.Println(c.Rc() + "Ryan wuz hear" + c.X)
	}
}

func main1() {
	//Prints "Hello World!" to the command line.
	fmt.Println(c.Clear + "Ryan wuz hear" + c.X)
}

func printMulti() {
	//Prints
	fmt.Println(c.Clear + c.Multi("Ryan wuz hear") + c.X)
}
