package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	consolereader := bufio.NewReader(os.Stdin)

	fmt.Println("I AM A PROGRAMMER!")

	fmt.Println("ARE YOU?")

	input, err := consolereader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if input == ("yes") {
		fmt.Println(input)
	}
}
