package main

import (
	"fmt"

	ext "github.com/test1/ext"
)

func main() {
	r := ext.Reader()
	fmt.Printf("hasil: %s", ext.Sorting(r))
}
