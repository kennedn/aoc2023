package main

import (
	"fmt"
	"testing"
)

func TestFiveOfAKind(t *testing.T) {
	testCases := []struct {
		s      string
		expect bool
	}{
		{
			s:      "AAAAA",
			expect: true,
		},
		{
			s:      "AABAA",
			expect: false,
		},
		{
			s:      "JJJJJ",
			expect: true,
		},
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
			res := fiveOfAKind(tc.s)
			if res != tc.expect {
				t.Errorf("Expected: %t Got: %t", tc.expect, res)
			}
		})
	}
}

func TestFourofAKind(t *testing.T) {
	testCases := []struct {
		s      string
		expect bool
	}{
		{
			s:      "AABAA",
			expect: true,
		},
		{
			s:      "BAAAA",
			expect: true,
		},
		{
			s:      "AAAAA",
			expect: false,
		},
		{
			s:      "ABABA",
			expect: false,
		},
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
			res := fourOfAKind(tc.s)
			if res != tc.expect {
				t.Errorf("Expected: %t Got: %t", tc.expect, res)
			}
		})
	}
}

func TestFullHouse(t *testing.T) {
	testCases := []struct {
		s      string
		expect bool
	}{
		{
			s:      "23332",
			expect: true,
		},
		{
			s:      "22223",
			expect: false,
		},
		{
			s:      "32T3K",
			expect: false,
		},
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
			res := fullHouse(tc.s)
			if res != tc.expect {
				t.Errorf("Expected: %t Got: %t", tc.expect, res)
			}
		})
	}
}

func TestThreeOfAKind(t *testing.T) {
	testCases := []struct {
		s      string
		expect bool
	}{
		{
			s:      "23332",
			expect: false,
		},
		{
			s:      "23331",
			expect: true,
		},
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
			res := threeOfAKind(tc.s)
			if res != tc.expect {
				t.Errorf("Expected: %t Got: %t", tc.expect, res)
			}
		})
	}
}

func TestTwoPair(t *testing.T) {
	testCases := []struct {
		s      string
		expect bool
	}{
		{
			s:      "23321",
			expect: true,
		},
		{
			s:      "23332",
			expect: false,
		},
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
			res := twoPair(tc.s)
			if res != tc.expect {
				t.Errorf("Expected: %t Got: %t", tc.expect, res)
			}
		})
	}
}

func TestOnePair(t *testing.T) {
	testCases := []struct {
		s      string
		expect bool
	}{
		{
			s:      "233A1",
			expect: true,
		},
		{
			s:      "23321",
			expect: false,
		},
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
			res := onePair(tc.s)
			if res != tc.expect {
				t.Errorf("Expected: %t Got: %t", tc.expect, res)
			}
		})
	}
}

func TestHighCard(t *testing.T) {
	testCases := []struct {
		s      string
		expect bool
	}{
		{
			s:      "12345",
			expect: true,
		},
		{
			s:      "12344",
			expect: false,
		},
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
			res := highCard(tc.s)
			if res != tc.expect {
				t.Errorf("Expected: %t Got: %t", tc.expect, res)
			}
		})
	}
}
