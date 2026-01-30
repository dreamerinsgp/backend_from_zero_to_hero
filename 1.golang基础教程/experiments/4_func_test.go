package experiments

import (
	"fmt"
	"testing"
)

func printHelloWorld(input string) string {
	return input
}

func TestFunc(t *testing.T) {
	// var input = "Hello,World!"
	// var response = printHelloWorld(input)
	input := "Hello,World!"
	response := printHelloWorld(input)

	fmt.Println(response)
}
