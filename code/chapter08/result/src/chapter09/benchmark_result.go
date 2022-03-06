package main

import (
	"bytes"
	"fmt"
	"testing"
	"text/template"
)

func main() {
	benchmarkResult := testing.Benchmark(func(b *testing.B) {
		templ := template.Must(template.New("test").Parse("Hello, {{.}}!"))
		// RunParallel will create GOMAXPROCS goroutines
		// and distribute work among them.
		b.RunParallel(func(pb *testing.PB) {
			// Each goroutine has its own bytes.Buffer.
			var buf bytes.Buffer
			for pb.Next() {
				// The loop body is executed b.N times total across all goroutines.
				buf.Reset()
				templ.Execute(&buf, "World")
			}
		})
	})

	// fmt.Printf("%8d\t%10d ns/op\t%10d B/op\t%10d allocs/op\n", benchmarkResult.N, benchmarkResult.NsPerOp(), benchmarkResult.AllocedBytesPerOp(), benchmarkResult.AllocsPerOp())
	fmt.Printf("%s\t%s\n", benchmarkResult.String(), benchmarkResult.MemString())
}
