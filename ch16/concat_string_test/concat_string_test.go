package concat_string_test

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"testing"
)

/*
goos: darwin
goarch: amd64
pkg: github.com/u8x9/go-learning/ch16/concat_string_test
BenchmarkByFmt-16              	   90073	     13165 ns/op
BenchmarkByByteBuffer-16       	 1329682	       915 ns/op
BenchmarkByStringBuilder-16    	 1615543	       752 ns/op
BenchmarkByAdd-16              	  247128	      4619 ns/op
PASS
ok  	github.com/u8x9/go-learning/ch16/concat_string_test	7.459s
*/

const numbers = 100

func BenchmarkByFmt(b *testing.B) {
	b.ResetTimer()
	for idx := 0; idx < b.N; idx++ {
		var s string
		for i := 0; i < numbers; i++ {
			s = fmt.Sprintf("%v%v", s, i)
		}
	}
	b.StopTimer()
}

func BenchmarkByByteBuffer(b *testing.B) {
	b.ResetTimer()
	for idx := 0; idx < b.N; idx++ {
		var buf bytes.Buffer
		for i := 0; i < numbers; i++ {
			buf.WriteString(strconv.Itoa(i))
		}
		buf.String()
	}
	b.StopTimer()
}
func BenchmarkByStringBuilder(b *testing.B) {
	b.ResetTimer()
	for idx := 0; idx < b.N; idx++ {
		var builder strings.Builder
		for i := 0; i < numbers; i++ {
			builder.WriteString(strconv.Itoa(i))
		}
		builder.String()
	}
	b.StopTimer()
}

func BenchmarkByAdd(b *testing.B) {
	b.ResetTimer()
	for idx := 0; idx < b.N; idx++ {
		var s string
		for i := 0; i < numbers; i++ {
			s += strconv.Itoa(i)
		}
	}
	b.StopTimer()
}
