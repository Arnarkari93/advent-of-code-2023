package main

import (
	"reflect"
	"testing"
)

func TestReadSeeds(t *testing.T) {
	input := "seeds: 79 14 55 13"
	result := ReadSeeds(input)
	expected := []int{79, 14, 55, 13}

if !reflect.DeepEqual(result, expected) {
		t.Errorf("Arrays are not equal: %v and %v", result, expected)
	}
}

