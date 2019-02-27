package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func sorting(w string) {
	s := strings.Split(w, "")
	sort.Strings(s)

	fmt.Println(strings.Join(s, ""))
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	sorting(text)
}
