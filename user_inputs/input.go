package input

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func GetSpins() int {
	fmt.Print("> ")
	reader := bufio.NewReader(os.Stdin)
	sentence, _ := reader.ReadString('\n')
	trimedSentence := strings.TrimSpace(sentence)
	bet, _ := strconv.Atoi(trimedSentence)
	return bet
}
// TODO: Need to add the errors that are return
func GetMoney() float64 {
	fmt.Print("> ")
	reader := bufio.NewReader(os.Stdin)
	sentence, _ := reader.ReadString('\n')
	trimedSentence := strings.TrimSpace(sentence)
	money, _ := strconv.ParseFloat(trimedSentence, 64)
	return money
}

func GetAnswer() bool {
	fmt.Println("Enter Yes or No")
	fmt.Print("> ")
	reader := bufio.NewReader(os.Stdin)
	reply, _ := reader.ReadString('\n')
	answer := strings.TrimSpace(strings.ToLower(reply))
	return strings.HasPrefix(answer, "y")
}
func GetMachineName() string {
	fmt.Println("Enter The Slot Machine Name.")
	fmt.Print("> ")
	reader := bufio.NewReader(os.Stdin)
	sentence, _ := reader.ReadString('\n')
	return strings.TrimSpace(sentence)
}

