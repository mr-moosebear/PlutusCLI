package main

import (
	//"encoding/json"
	"fmt"
	"input"
	//"os"
	"save_load"
)

/*type MachineData struct {
	MachineName string `json:"machineName"`
	StartingMoney float64 `json:"startingMoney"`
	Bet float64 `json:"bet"`
	Spins int `json:"spins"`
}

func save(data MachineData) {
	file, _ := os.Create("hello.json")
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") //NOTE: Might need work or need gone
	encode := encoder.Encode(data)
	if encode != nil {
		fmt.Println("hmmm")
	}
}*/

func main() {
	fmt.Println("Welcome to Plutus Command Line Interface!")
	name := input.GetMachineName()
	fmt.Println("Enter your starting money.")
	money := input.GetMoney()
	fmt.Println("Enter the amount you would like to bet per spin.")
	bet := input.GetMoney()
	fmt.Println("Enter Amount of Spins you would like to do.")
	spins := input.GetSpins()

	info := save_load.MachineData{ MachineName: name, StartingMoney: money, Bet: bet, Spins: spins }
	//save(info)
	save_load.Save(info)
	fmt.Println(info.MachineName, info.StartingMoney, info.Bet, info.Spins)
}
