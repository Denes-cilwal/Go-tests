package go_test_demo

import "testing"

// multiple test func inside single test file calle test suite
func TestHelloEmptyArgs(t *testing.T) {

	// test - for empty argument
	// define expected result
	emptyResult := Hello("") // should return "Hello Dude!"

	// result of execution not valid call
	if emptyResult != "Hello Dude!" {
		t.Errorf("hello failed,expected %v, got %v", "Hello Dude!", emptyResult)
	}

	// test for valid argument
	result := Hello("Mike") // should return Mike
	if result != "Hello Mike!" {
		t.Errorf("hello failed,expected %v, got %v", "Hello Dude!", result)
	} else {
		t.Logf("hello success,expected %v, got %v", "Hello Dude!", result)
	}
}

func TestHelloValidArgs(t *testing.T) {

	// test for valid argument
	result := Hello("Mike") // should return Mike
	if result != "Hello Mike!" {
		t.Errorf("hello failed,expected %v, got %v", "Hello Mike!", result)
	} else {
		t.Logf("hello success,expected %v, got %v", "Hello Mike!", result)

	}
}
