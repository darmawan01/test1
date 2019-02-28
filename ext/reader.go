package ext

import (
	"bufio"
	"fmt"
	"os"
)

//Reader for read input
func Reader() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')

	return text
}
