package main

import (
	"fmt"
	"input"
	"os"
	"save_load"
	"strings"
)

/*type MachineData struct {
	MachineName string `json:"machineName"`
	StartingMoney float64 `json:"startingMoney"`
	Bet float64 `json:"bet"`
	Spins int `json:"spins"`
}*/

func getFileNames() []string {
	files, _ := os.ReadDir("data")
	var fileNames []string
	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}
	return fileNames

}
func displayFileNames(files []string) {
	for _, file := range files {
		fmt.Println(strings.TrimSuffix(file, ".json"))
	}
	fmt.Println("What File would you like to load???")
}

func main() {
	fmt.Println("Welcome to Plutus Command Line Interface!")
	fmt.Println("Would you like to start a new machine?")
	answer := input.GetAnswer()
	if answer == true {
		files := getFileNames()
		displayFileNames(files)
	} else {
		name := input.GetMachineName()
		fmt.Println("Enter your starting money.")
		money := input.GetMoney()
		fmt.Println("Enter the amount you would like to bet per spin.")
		bet := input.GetMoney()
		fmt.Println("Enter Amount of Spins you would like to do.")
		spins := input.GetSpins()

		info := save_load.MachineData{ MachineName: name, StartingMoney: money, Bet: bet, Spins: spins }
		save_load.Save(info)
		t, _ := save_load.LoadFile()
		fmt.Println("Saved file is ", info.MachineName, info.StartingMoney, info.Bet, info.Spins)
		fmt.Println("Load file struct is", t, t.Bet)
	}
}
