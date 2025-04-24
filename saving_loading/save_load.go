package save_load

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"os"
)

type MachineData struct {
	MachineName string `json:"machineName"`
	StartingMoney float64 `json:"startingMoney"`
	Bet float64 `json:"bet"`
	Spins int `json:"spins"`
	TotalMoneyAdded float64 `json:"totalMoneyAdded"`
}

func LoadFile() (MachineData, error) {
	var fileData MachineData
	filepath := filepath.Join("data", "amount.json")
	file, err := os.Open(filepath)
	if err != nil {
		return fileData, err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&fileData)
	return fileData, err
}
func createFileName(str string) {
	
}
func Save(data MachineData) {
	filepath := filepath.Join("data", "amount.json")
	//fmt.Println(filePath)
	file, _ := os.Create(filepath)
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") //NOTE: Might need work or need gone
	encode := encoder.Encode(data)
	if encode != nil {
		fmt.Println("hmmm")
	}
}
