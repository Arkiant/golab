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
	Price      float32
	Name       string
	Surname    string
	Parameters []Parameters
}

func serialize(item SerializeStruct, sep, sepLevel2, paramSep string) string {
	const (
		name int = iota
		surname
		price
		parameters
	)

	const (
		paramKey int = iota
		paramValue
	)

	b := bytes.Buffer{}
	b.WriteString(item.Name)
	b.WriteString(sep)
	b.WriteString(item.Surname)
	b.WriteString(sep)
	b.WriteString(strconv.FormatFloat(float64(item.Price), 'f', 2, 32))
	b.WriteString(sep)
	for i, p := range item.Parameters {
		b.WriteString(p.Key)
		b.WriteString(sepLevel2)
		b.WriteString(p.Value)
		if i < len(item.Parameters)-1 {
			b.WriteString(paramSep)
		}
	}

	return b.String()

}
