package token

import (
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
// BenchmarkSerialize-8   	 2000000	      899 ns/op	    1579 B/op	       0 allocs/op
// PASS
// ok  	github.com/arkiant/golab/token	3.396s
// Success: Benchmarks passed.

func serialize(buf *[]byte, item SerializeStruct, sep, sepLevel2, paramSep string) []byte {

	*buf = append(*buf, item.Name...)
	*buf = append(*buf, sep...)
	*buf = append(*buf, item.Surname...)
	*buf = append(*buf, sep...)
	*buf = strconv.AppendFloat(*buf, item.Price, 'f', -1, 64)
	*buf = append(*buf, sep...)

	for i, p := range item.Parameters {
		*buf = append(*buf, p.Key...)
		*buf = append(*buf, sepLevel2...)
		*buf = append(*buf, p.Value...)

		if i < len(item.Parameters)-1 {
			*buf = append(*buf, paramSep...)
		}
	}

	return *buf

}
