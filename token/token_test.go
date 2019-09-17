package token

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSerialize(t *testing.T) {

	t.Parallel()

	tc := []struct {
		Name      string
		In        SerializeStruct
		Out       []byte
		Sep       string
		SepLevel2 string
		ParamSep  string
	}{
		{
			"Serialize struct to string",
			SerializeStruct{
				Name:    "test",
				Surname: "testSurname",
				Price:   16.00,
				Parameters: []Parameters{
					{Key: "TestParameter", Value: "ValueParameter"},
					{Key: "TestParameter2", Value: "ValueParameter2"},
					{Key: "TestParameter3", Value: "ValueParameter3"},
				},
			},
			[]byte("test#testSurname#16#TestParameter##ValueParameter###TestParameter2##ValueParameter2###TestParameter3##ValueParameter3"),
			"#",
			"##",
			"###",
		},
		{
			"Serialize struct to string",
			SerializeStruct{
				Name:    "test",
				Surname: "testSurname",
				Price:   24.22,
				Parameters: []Parameters{
					{Key: "TestParameter", Value: "ValueParameter"},
					{Key: "TestParameter2", Value: "ValueParameter2"},
					{Key: "TestParameter3", Value: "ValueParameter3"},
				},
			},
			[]byte("test#testSurname#24.22#TestParameter##ValueParameter###TestParameter2##ValueParameter2###TestParameter3##ValueParameter3"),
			"#",
			"##",
			"###",
		},
		{
			"Serialize struct to string",
			SerializeStruct{
				Name:    "test",
				Surname: "testSurname",
				Price:   21.50,
				Parameters: []Parameters{
					{Key: "TestParameter", Value: "ValueParameter"},
					{Key: "TestParameter2", Value: "ValueParameter2"},
					{Key: "TestParameter3", Value: "ValueParameter3"},
				},
			},
			[]byte("test#testSurname#21.5#TestParameter##ValueParameter###TestParameter2##ValueParameter2###TestParameter3##ValueParameter3"),
			"#",
			"##",
			"###",
		},
	}

	for _, tt := range tc {
		t.Run(tt.Name, func(t *testing.T) {
			buf := make([]byte, 0, 0)
			s := serialize(&buf, tt.In, tt.Sep, tt.SepLevel2, tt.ParamSep)
			assert.Equal(t, tt.Out, s)
		})
	}

}

func BenchmarkSerialize(b *testing.B) {

	s := SerializeStruct{
		Name:    "test",
		Surname: "testSurname",
		Price:   16.0,
		Parameters: []Parameters{
			{Key: "TestParameter", Value: "ValueParameter"},
			{Key: "TestParameter2", Value: "ValueParameter2"},
			{Key: "TestParameter3", Value: "ValueParameter3"},
			{Key: "TestParameter3", Value: "ValueParameter3"},
			{Key: "TestParameter3", Value: "ValueParameter3"},
			{Key: "TestParameter3", Value: "ValueParameter3"},
			{Key: "TestParameter3", Value: "ValueParameter3"},
		},
	}

	buf := make([]byte, 0, 1000)
	b.ResetTimer()
	b.SetParallelism(4)
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		serialize(&buf, s, "#", "##", "###")
	}
}

func TestDeserialize(t *testing.T) {
	tc := []struct {
		Name string
		In   string
		Out  SerializeStruct
	}{
		{
			"Deserialize string to struct",
			"test#testSurname#16.0#TestParameter[ValueParameter",
			SerializeStruct{
				Name:    "test",
				Surname: "testSurname",
				Price:   16.0,
				Parameters: []Parameters{
					{Key: "TestParameter", Value: "ValueParameter"},
				},
			},
		},
	}

	for _, tt := range tc {
		t.Run(tt.Name, func(t *testing.T) {

		})
	}
}
