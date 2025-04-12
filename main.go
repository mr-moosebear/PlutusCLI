package main

import (
	"fmt"
	"input"
)

func main() {
	fmt.Println("Welcome to Plutus Command Line Interface!")
	val := input.GetMachineName()
	fmt.Println(val)
}
