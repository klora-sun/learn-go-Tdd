package reapt

import "testing"

func TestRepeat(t *testing.T) {
	test := Repeat("a",5)
	expected := "aaaaa"
	if test != expected {
		t.Errorf("expected '%s' but got '%s'", expected, test)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

