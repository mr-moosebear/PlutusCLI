package input

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)
// TODO: Need to add the errors that are return
func GetMoney() float64 {
	fmt.Print("> ")
	reader := bufio.NewReader(os.Stdin)
	sentence, _ := reader.ReadString('\n')
	trimedSentence := strings.TrimSpace(sentence)
	money, _ := strconv.ParseFloat(trimedSentence, 64)
	return money
}

func GetMachineName() string {
	fmt.Println("Enter The Slot Machine Name.")
	fmt.Print("> ")
	reader := bufio.NewReader(os.Stdin)
	sentence, _ := reader.ReadString('\n')
	return strings.TrimSpace(sentence)
}

