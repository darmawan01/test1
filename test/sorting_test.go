package test

import (
	"testing"

	ext "github.com/test1/ext"
)

func TestResult(t *testing.T) {
	s := "awan"
	result := ext.Sorting(s)

	if result != "aanw" {
		t.Errorf("Result invalid, should be 'aanw' from word 'awan'")
	}
}
