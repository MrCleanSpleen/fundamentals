package main

import (
	"fmt"
)

func main() {
	var myname string
	str1 := "What is your name?"
	fmt.Println(str1)
	fmt.Scanf("%s", &myname)
	fmt.Println("Hello", myname)
}
