package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func sorting(s string) {
	text := strings.Split(s, "")
	sort.Strings(text)

	fmt.Printf("hasil: %s", strings.Trim(strings.Join(text, ""), "\n"))
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')

	sorting(text)
}
