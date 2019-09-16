package token

import (
	"bytes"
	"strconv"
)

type Parameters struct {
	Key   string
	Value string
}

type SerializeStruct struct {
	Price      float64
	Name       string
	Surname    string
	Parameters []Parameters
}

// Benchmark about this implementation
//
// goos: windows
// goarch: amd64
// pkg: github.com/arkiant/golab/token
// BenchmarkSerialize-8   	 2000000	       636 ns/op	     562 B/op	       5 allocs/op
// PASS
// ok  	github.com/arkiant/golab/token	2.265s
// Success: Benchmarks passed.

func serialize(item SerializeStruct, sep, sepLevel2, paramSep string) []byte {
	b := bytes.Buffer{}

	b.WriteString(item.Name)
	b.WriteString(sep)
	b.WriteString(item.Surname)
	b.WriteString(sep)
	b.WriteString(strconv.FormatFloat(item.Price, 'f', -1, 64))
	b.WriteString(sep)

	for i, p := range item.Parameters {
		b.WriteString(p.Key)
		b.WriteString(sepLevel2)
		b.WriteString(p.Value)

		if i < len(item.Parameters)-1 {
			b.WriteString(paramSep)
		}
	}

	return b.Bytes()

}
