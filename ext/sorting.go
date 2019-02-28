package ext

import (
	"sort"
	"strings"
)

// Sorting handle sort character
func Sorting(s string) string {
	text := strings.Split(s, "")
	sort.Strings(text)

	return strings.Trim(strings.Join(text, ""), "\n")
}
