package main

import (
	h "hello"
	"fmt"
	c "github.com/skilstak/go/colors"
	//"strings"
	//i "github.com/skilstak/go/input"
	"os"
)

func main() {
	name, option = h.ParseArgs()
	message := "Hello " + name + "!"

	if option == "-m" {
		h.PrintMulti()
	} else if option == "-r" {
		h.PrintRandom()
	} else if option == "-f" {

