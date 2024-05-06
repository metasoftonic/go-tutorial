package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("Test character repeat", func(t *testing.T) {
		got := Repeat("a", 5)
		expected := "aaaaa"

		if got != expected {
			t.Errorf("Expected %q, got %q", expected, got)
		}
	})
}
func ExampleRepeat() {
	repeated := Repeat("b", 4)
	fmt.Println(repeated)
	//Output: bbbb
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 6)
	}
}
