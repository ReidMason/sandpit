package main

import (
	"log"
)

func main() {
	b := byte(0b1101011)
	b2 := byte(0b1110011)

	// binary := strconv.FormatInt(n, 2)
	printByte(b)
	printByte(b2)
	printByte(b & b2)
	printByte(b & b2)
	printByte(b | b2)
}

func printByte(b byte) {
	log.Printf("%08b", b)
}
