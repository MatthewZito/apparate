package internal

import (
	"errors"
	"fmt"
	"os"
)

// Itob, given an integer, returns a boolean value
func Itob(n int) bool {
	if n > 0 {
		return true
	} else {
		return false
	}
}

// Exists determines whether a given file path is extant on the host system
func Exists(path string) bool {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) || err != nil {
		return false
	}

	return true
}

// ErrExit prints an error message to stderr and exits with an erroneous return code
func ErrExit(m string) {
	println(m)
	os.Exit(1)
}

// OkExit prints a notify message to stdout and exits with a 0 return code
func OkExit(m string) {
	fmt.Println("Apparate: " + m)
	os.Exit(0)
}
