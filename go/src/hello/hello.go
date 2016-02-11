package hello

import (
	"fmt"
	c "github.com/skilstak/go/colors"
	s "strings"
	//i "github.com/skilstak/go/input"
	"os"
	"time"
)

func PrintPlain(message string) {
	//Prints "Hello World!" to the command line.
	fmt.Println(c.Clear + "Hello " + message + c.X)
}
func PrintRandom(message string) {
	//Prints "Hello World" to the command line in random color.
	fmt.Println(c.Clear + c.Rc() + "Hello " + message + c.X)
}
func PrintMulti(message string) {
	//Prints Hello World to the command line with each character a different color.
	for {
		fmt.Println(c.Multi("Hello " + message + c.Clear))
		time.Sleep(50000)
		fmt.Print(c.Clear)
	}
}
func PrintForevah(message string) {
	//Prints "Hello World" in a random color forever in a loop.
	for {
		fmt.Print(c.Rc() + "Hello " + message + c.X)
	}
}

func ParseArgs() (string, string) {
	//Parses arguments from the command line
	name := "YOU"
	option := ""
	if len(os.Args) > 2 {
		option = os.Args[1]
		name = os.Args[2]
	}
	if len(os.Args) == 2 {
		if s.HasPrefix(os.Args[1], "-") {
			option = os.Args[1]
		} else {
			name = os.Args[1]
		}
	}
	return name, option
}
