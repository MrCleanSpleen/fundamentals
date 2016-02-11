package main

import (
	h "hello"
)

func main() {
	name, option := h.ParseArgs()
	message := "Hello " + name + "!"

	if option == "-m" {
		h.PrintMulti()
	} else if option == "-r" {
		h.PrintRandom()
	} else if option == "-f" {
		h.PrintForevah()
	} else {
		h.PrintPlain()
	}
}
