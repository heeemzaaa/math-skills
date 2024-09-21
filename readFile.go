package ms

import (
	"os"
	"strings"
)

// This function read the givven file and return every number in a string inside a slice of strings
func ReadFile(s string) []string {
	file, err := os.ReadFile(s)
	if err != nil {
		return nil
	}
	SplitFile := strings.Fields(string(file))
	return SplitFile
}
