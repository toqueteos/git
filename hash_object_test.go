package git

import (
	"testing"
)

var hashObjectTests = []struct {
	in, out string
}{
	{"test content", "d670460b4b4aece5915caf5c68d12f560a9fe3e4"},
}

func TestHashObject(t *testing.T) {
	for index, tt := range hashObjectTests {
		s, err := HashObject(tt.in, false)
		if err != nil {
			t.Fatalf("%d. %q failed: %s\n", index, tt.in, err)
		}

		if s != tt.out {
			t.Fatalf("%d. HashObject(%q) != %q, got %s.\n", index, tt.in, tt.out, s)
		}
	}
}
