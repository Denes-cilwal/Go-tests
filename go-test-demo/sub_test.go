package go_test_demo

import "testing"

func TestSub(t *testing.T) {
	x, y := 6, 3
	actualResult := sub(x, y)
	expectedResult := 4

	if actualResult != expectedResult {
		t.Fail()
	}

}
