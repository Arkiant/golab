package token

import "testing"

type Parameters struct {
	Key   string
	Value string
}

type SerializeStruct struct {
	Name       string
	Surname    string
	Price      float32
	Parameters []Parameters
}

func serializeTest(t *testing.T) {

	tc := []struct {
		Name string
		In   SerializeStruct
		Out  string
	}{
		{
			"Serialize struct to string",
			SerializeStruct{
				Name:    "test",
				Surname: "testSurname",
				Price:   16.0,
				Parameters: []Parameters{
					{Key: "TestParameter", Value: "ValueParameter"},
				},
			},
			"test#testSurname#16.0#TestParameter[ValueParameter",
		},
	}

	for _, tt := range tc {
		t.Run(tt.Name, func(t *testing.T) {

		})
	}

}

func deserializeTest(t *testing.T) {
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
