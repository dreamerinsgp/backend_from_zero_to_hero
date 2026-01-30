package experiments

import (
	"fmt"
	"testing"
)

func vals() (int, int) {
	return 2, 3
}

func TestFuncMultipleResponse(t *testing.T) {
	response1, reponse2 := vals()

	fmt.Println(response1, reponse2)
}
