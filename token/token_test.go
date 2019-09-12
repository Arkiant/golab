package token

import "testing"

func serializeTest(t *testing.T) {

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

	tc := []struct {
		Name     string
		ItemTest SerializeStruct
	}{
		{"Serialize struct test", SerializeStruct{
			Name:    "test",
			Surname: "testSurname",
			Price:   16.0,
			Parameters: []Parameters{
				{Key: "TestParameter", Value: "ValueParameter"},
			},
		}},
	}

	for _, tt := range tc {
		t.Run(tt.Name, func(t *testing.T) {

		})
	}

}

func deserializeTest(t *testing.T) {

}
