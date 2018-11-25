package lib

import (
	"fmt"
	"testing"
)

func stringPtr(s string) *string {
	return &s
}

func stringNilPtr() *string {
	return nil
}

var isNilTests = []struct {
	in  interface{}
	out bool
}{
	{nil, true},
	{stringNilPtr(), true},
	{"a", false},
	{stringPtr("a"), false},
}

func TestIsNil(t *testing.T) {
	for i, c := range isNilTests {
		t.Run(fmt.Sprintf("TestIsNil %d", i), func(t *testing.T) {
			if IsNil(c.in) != c.out {
				t.Errorf("got %v, want %v", IsNil(c.in), c.out)
			}
		})
	}
}

var isZeroOrNilTests = []struct {
	in  interface{}
	out bool
}{
	{nil, true},
	{stringNilPtr(), true},
	{"a", false},
	{stringPtr("a"), false},
	{stringPtr(""), false},
	{"", true},
}

func TestIsZeroOrNil(t *testing.T) {
	for i, c := range isZeroOrNilTests {
		t.Run(fmt.Sprintf("TestIsZeroOrNil %d", i), func(t *testing.T) {
			if IsZeroOrNil(c.in) != c.out {
				t.Errorf("got %v, want %v", IsZeroOrNil(c.in), c.out)
			}
		})
	}
}

var testValueTests = []struct {
	in  interface{}
	out interface{}
}{
	{"a", "a"},
	{stringPtr("a"), "a"},
	{stringPtr(""), ""},
	{"", ""},
}

func TestValue(t *testing.T) {
	for i, c := range testValueTests {
		t.Run(fmt.Sprintf("TestValue %d", i), func(t *testing.T) {
			if Value(c.in) != c.out {
				t.Errorf("got %v, want %v", Value(c.in), c.out)
			}
		})
	}
}
