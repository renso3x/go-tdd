package integers

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {

	t.Run("can add 2 + 2", func(t *testing.T) {
		sum := add(2, 2)
		expected := 4

		if sum != expected {
			t.Errorf("expected '%d' but got '%d'", expected, sum)
		}
	})
}

func ExampleAdder() {
	sum := add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
