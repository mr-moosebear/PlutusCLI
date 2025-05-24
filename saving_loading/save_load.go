package save_load

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"os"
	"strings"
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
	t := getFileNames()
	fmt.Println("File Names are ", t)
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
func getFileNames() []string {
	files, _ := os.ReadDir("data")
	var fileNames []string
	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}
	return fileNames

}
func createFileName(str string) string {
	splitStrings := strings.Split(str, " ")
	name := strings.Join(splitStrings, "")
	fileNames := []string{name, ".json"}
	fileName := strings.Join(fileNames, "")
	return fileName
}
func Save(data MachineData) {
	name := createFileName(data.MachineName)
	filepath := filepath.Join("data", name)
	file, _ := os.Create(filepath)
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") //NOTE: Might need work or need gone
	encode := encoder.Encode(data)
	if encode != nil {
		fmt.Println("hmmm")
	}
}
