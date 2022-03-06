package testing_test

import (
	"bytes"
	"html/template"
	"testing"
	"time"

	. "chapter09/testing"
)

// Table-Driven Test
func TestFib_TableDrivenParallel(t *testing.T) {
	var fibTests = []struct {
		name     string
		in       int // input
		expected int // expected result
	}{
		{"1的Fib", 1, 1},
		{"2的Fib", 2, 1},
		{"3的Fib", 3, 2},
		{"4的Fib", 4, 3},
		{"5的Fib", 5, 5},
		{"6的Fib", 6, 8},
		{"7的Fib", 7, 13},
	}

	for _, tt := range fibTests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Log("time:", time.Now())
			t.Parallel()
			time.Sleep(3 * time.Second)
			actual := Fib(tt.in)
			if actual != tt.expected {
				t.Errorf("Fib(%d) = %d; expected %d", tt.in, actual, tt.expected)
			}
		})
	}

	for _, tt := range fibTests {
		t.Log("time:", time.Now())
		actual := Fib(tt.in)
		if actual != tt.expected {
			t.Errorf("Fib(%d) = %d; expected %d", tt.in, actual, tt.expected)
		}
	}

	defer func() {
		t.Log("time:", time.Now())
	}()
}

func TestFib(t *testing.T) {
	var (
		in       = 7  // input
		expected = 13 // expected result
	)
	actual := Fib(in)
	if actual != expected {
		t.Errorf("Fib(%d) = %d; expected %d", in, actual, expected)
	}
}

// Table-Driven Test
func TestFib_TableDriven(t *testing.T) {
	var fibTests = []struct {
		in       int // input
		expected int // expected result
	}{
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
	}

	for _, tt := range fibTests {
		actual := Fib(tt.in)
		if actual != tt.expected {
			t.Errorf("Fib(%d) = %d; expected %d", tt.in, actual, tt.expected)
		}
	}
}

func BenchmarkFib1(b *testing.B)  { benchmarkFib(1, b) }
func BenchmarkFib2(b *testing.B)  { benchmarkFib(2, b) }
func BenchmarkFib3(b *testing.B)  { benchmarkFib(3, b) }
func BenchmarkFib10(b *testing.B) { benchmarkFib(10, b) }
func BenchmarkFib20(b *testing.B) { benchmarkFib(20, b) }
func BenchmarkFib40(b *testing.B) { benchmarkFib(40, b) }

func benchmarkFib(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fib(i)
	}
}

func BenchmarkTmplExucte(b *testing.B) {
	b.ReportAllocs()

	templ := template.Must(template.New("test").Parse("Hello, {{.}}!"))
	b.RunParallel(func(pb *testing.PB) {
		// Each goroutine has its own bytes.Buffer.
		var buf bytes.Buffer
		for pb.Next() {
			// The loop body is executed b.N times total across all goroutines.
			buf.Reset()
			templ.Execute(&buf, "World")
		}
	})
}
