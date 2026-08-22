package internal

import (
	"testing"
)

func TestHashSet(t *testing.T) {
	hs := NewEmptyHashset[string]()

	if hs.Len() > 0 {
		t.Fatal()
	}

	hs.Set("hello")
	hs.Set("hello")
	hs.Set("hello")
	hs.Set("hello")
	if hs.Len() != 1 {
		t.Fatal()
	}
	if !hs.Exists("hello") {
		t.Fatal()
	}

	hs.Set("hell")
	hs.Set("hell")
	hs.Set("hell")
	if hs.Len() != 2 {
		t.Fatal()
	}
	if !hs.Exists("hell") {
		t.Fatal()
	}

	hs.Remove("hello")
	if hs.Len() != 1 {
		t.Fatal()
	}
	if hs.Exists("hello") {
		t.Fatal()
	}

	hs.Remove("hell")
	if hs.Len() != 0 {
		t.Fatal()
	}
	if hs.Exists("hell") {
		t.Fatal()
	}
}
