package benchmark_test

/*
 $ go test -bench=. -benchmem
goos: darwin
goarch: amd64
pkg: github.com/u8x9/go-learning/ch11/benchmark_test
BenchmarkConcatStringByAdd-16            	10021892	       113 ns/op	      16 B/op	       4 allocs/op
BenchmarkConcatStringByBytesBuffer-16    	20176182	        59.1 ns/op	      64 B/op	       1 allocs/op
BenchmarkConcatStringBySlice-16          	 9209010	       115 ns/op	      85 B/op	       2 allocs/op
PASS
ok  	github.com/u8x9/go-learning/ch11/benchmark_test	3.832s
*/

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConcatStringByAdd(t *testing.T) {
	assert := assert.New(t)
	elems := []string{"1", "2", "3", "4", "5"}
	ret := ""
	for _, elem := range elems {
		ret += elem
	}
	assert.Equal("12345", ret)
}

func BenchmarkConcatStringByAdd(b *testing.B) {
	elems := []string{"1", "2", "3", "4", "5"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ret := ""
		for _, elem := range elems {
			ret += elem
		}
	}
	b.StopTimer()
}

func TestConcatStringByBytesBuffer(t *testing.T) {
	assert := assert.New(t)
	var buf bytes.Buffer
	elems := []string{"1", "2", "3", "4", "5"}
	for _, elem := range elems {
		buf.WriteString(elem)
	}
	assert.Equal("12345", buf.String())
}

func BenchmarkConcatStringByBytesBuffer(b *testing.B) {
	elems := []string{"1", "2", "3", "4", "5"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		for _, elem := range elems {
			buf.WriteString(elem)
		}
		buf.String()
	}
	b.StopTimer()
}

func TestConcatStringBySlice(t *testing.T) {
	assert := assert.New(t)
	elems := []string{"1", "2", "3", "4", "5"}
	slice := make([]string, 0, len(elems))
	for _, elem := range elems {
		slice = append(slice, elem)
	}
	assert.Equal("12345", strings.Join(slice, ""))
}

func BenchmarkConcatStringBySlice(b *testing.B) {
	elems := []string{"1", "2", "3", "4", "5"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		slice := make([]string, 0, len(elems))
		for _, elem := range elems {
			slice = append(slice, elem)
		}
		strings.Join(slice, "")
	}
	b.StopTimer()
}
