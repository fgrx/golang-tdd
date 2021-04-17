package iterations

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("A")
	expected :="AAAAA"

	if repeated!=expected{
		t.Errorf("Expected %q but got %q",expected, repeated)
	}
}