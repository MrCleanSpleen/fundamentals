package main

import (
	h "hello"
)

func main() {
	name, option := h.ParseArgs()
	message := name + "!"

	if option == "-m" {
		h.PrintMulti(message)
	} else if option == "-r" {
		h.PrintRandom(message)
	} else if option == "-f" {
		h.PrintForevah(message)
	} else {
		h.PrintPlain(message)
	}
}
