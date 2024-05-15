package adder_test

import (
	"fmt"
	adder "learn_testing"
	"testing"
)

func TestAdd(t *testing.T) {
	want := 6
	got := adder.Add(2, 4)
	if want != got {
		t.Errorf("want %d, got %d\n", want, got)
	}
}

func BenchmarkTestAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		want := 6
		got := adder.Add(2, 4)
		if want != got {
			b.Errorf("want %d, got %d\n", want, got)
		}
	}
}

func FuzzAdd(f *testing.F) {
	for i := range 2 {
		f.Add(2, i)
	}

	f.Fuzz(func(t *testing.T, a, b int) {
		r1 := adder.Add(a, b)
		want := "int"
		got := fmt.Sprintf("%T", r1)
		if want != got {
			t.Errorf("types don't match. want %s, got %s", want, got)
		}

	})
}
