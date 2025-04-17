package main

import (
	"encoding/json"
	"fmt"
	"input"
	"os"
)

type MachineData struct {
	machineName string
	startingMoney float64
}

func save(data MachineData) {
	file, _ := os.Create("amount.json")
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") //NOTE: Might need work or need gone
	encode := encoder.Encode(data)
	if encode != nil {
		fmt.Println("hmmm")
	}
}
func main() {
	fmt.Println("Welcome to Plutus Command Line Interface!")
	name := input.GetMachineName()
	fmt.Println("Enter your starting money.")
	money := input.GetMoney()


	info := MachineData{ name, money }
	save(info)
	fmt.Println(info.machineName, info.startingMoney)
}
