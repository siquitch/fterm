package utils

import "testing"

// Makes sure a navigator object functions correctly
func TestNavigator(t *testing.T) {
	navigator := NewNavigator(0, 3)

	if navigator.Index() != 0 {
		t.Fail()
	}

	if navigator.length != 3 {
		t.Fail()
	}

	if navigator.ShouldShowHelp() {
		t.Fail()
	}

	navigator.Next()

	if navigator.Index() != 1 {
		t.Fail()
	}

	navigator.Previous()
	navigator.Previous()

	if navigator.Index() != 2 {
		t.Fail()
	}

	navigator.Reset(10)

	if navigator.Index() != 0 {
		t.Fail()
	}

	if navigator.length != 10 {
		t.Fail()
	}
}
