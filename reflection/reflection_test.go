package main

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one string field",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var actual []string
			walk(test.Input, func(input string) {
				actual = append(actual, input)
			})

			if !reflect.DeepEqual(actual, test.ExpectedCalls) {
				t.Errorf("Expected %v but got %v", test.ExpectedCalls, actual)
			}
		})
	}
}
