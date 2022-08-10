package iteration

import (
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {

	t.Run("run 5x times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})

	t.Run("run X times", func(t *testing.T) {
		repeated := Repeat("a", 20)
		expected := strings.Repeat("a", 20)

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
