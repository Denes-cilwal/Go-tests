package go_test_demo

import "testing"

// multiple test func inside single test file called test suite
func TestSayHelloEmptyArgs(t *testing.T) {

	// test - for empty argument
	// define expected result
	emptyResult := sayHello("") // should return "Hello Dude!"

	// result of execution not valid call
	if emptyResult != "Hello Anonymous!" {
		t.Errorf("sayHello failed,expected %v, got %v", "Hello Anonymous!", emptyResult)
	}

	// test for valid argument
	result := sayHello("Mike") // should return Mike
	if result != "Hello Mike" {
		t.Errorf("hello failed,expected %v, got %v", "Hello Anonymous!", result)
	} else {
		// provide non-failing information
		t.Logf("hello success,expected %v, got %v", "Hello Mike", result)
	}
}

func TestSayHelloValidArgs(t *testing.T) {

	// test for valid argument
	result := sayHello("Mike") // should return Mike
	if result != "Hello Mike" {
		t.Errorf("hello failed,expected %v, got %v", "Hello Mike", result)
	} else {
		t.Logf("hello success,expected %v, got %v", "Hello Mike", result)

	}
}

func TestSayHelloInvalidArgs(t *testing.T) {
	// test for invalid args

	result := sayHello("David")
	if result != "Hello David!" {
		t.Errorf("hello failed, expected %v, got %v", "Hello David!", result)
	} else {
		t.Logf("hello success,expected %v, got %v", "Hello David!", result)

	}
}
