package testing_test

import (
	"bytes"
	. "chapter09/testing"
	"html/template"
	"testing"
)

var pairs = []struct {
	k string
	v string
}{
	{"polaris", "徐新华"},
	{"studygolang", "Go语言中文网"},
	{"stdlib", "Go语言标准库"},
	{"polaris1", "徐新华1"},
	{"studygolang1", "Go语言中文网1"},
	{"stdlib1", "Go语言标准库1"},
	{"polaris2", "徐新华2"},
	{"studygolang2", "Go语言中文网2"},
	{"stdlib2", "Go语言标准库2"},
	{"polaris3", "徐新华3"},
	{"studygolang3", "Go语言中文网3"},
	{"stdlib3", "Go语言标准库3"},
	{"polaris4", "徐新华4"},
	{"studygolang4", "Go语言中文网4"},
	{"stdlib4", "Go语言标准库4"},
}

// 注意 TestWriteToMap 需要在 TestReadFromMap 之前
func TestWriteToMap(t *testing.T) {
	t.Parallel()
	for _, tt := range pairs {
		WriteToMap(tt.k, tt.v)
	}
}

func TestReadFromMap(t *testing.T) {
	t.Parallel()
	for _, tt := range pairs {
		actual := ReadFromMap(tt.k)
		if actual != tt.v {
			t.Errorf("the value of key(%s) is %s, expected: %s", tt.k, actual, tt.v)
		}
	}
}

func BenchmarkTemplateParallel(b *testing.B) {
	templ := template.Must(template.New("test").Parse("Hello, {{.}}!"))
	b.RunParallel(func(pb *testing.PB) {
		var buf bytes.Buffer
		for pb.Next() {
			buf.Reset()
			templ.Execute(&buf, "World")
		}
	})
}
