package input

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetStartingMoney() float64 {
	fmt.Println("Enter Your Starting Money.")
	fmt.Print("> ")
	
}

func GetMachineName() string {
	fmt.Println("Enter The Slot Machine Name.")
	fmt.Print("> ")
	reader := bufio.NewReader(os.Stdin)
	sentence, _ := reader.ReadString('\n')
	return strings.TrimSpace(sentence)
}

