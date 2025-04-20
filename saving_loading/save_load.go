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
}

func Save(data MachineData) {
	filePath := filepath.Join("data", "amount.json")
	fmt.Println(filePath)
	file, _ := os.Create(filePath)
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") //NOTE: Might need work or need gone
	encode := encoder.Encode(data)
	if encode != nil {
		fmt.Println("hmmm")
	}
}
