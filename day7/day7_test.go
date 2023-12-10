package main

import "testing"

func TestGetHandStrength(t *testing.T) {
	testCases := []struct {
		hand     string
		strenght HandStrength
	}{
		{"AAAAA", FiveOfAKind},
		{"AA8AA", FourOfAKind},
		{"23332", FullHouse},
		{"TTT98", ThreeOfAKind},
		{"23432", TwoPair},
		{"A23A4", OnePair},
		{"23456", HighCard},
	}
	for _, tc := range testCases {
		t.Run(tc.hand, func(t *testing.T) {
			result := GetHandStrength(tc.hand)
			if result != tc.strenght {
				t.Errorf("%s Expected %d, but got %d", tc.hand, tc.strenght, result)
			}
		})
	}

}
