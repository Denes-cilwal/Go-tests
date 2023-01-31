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

/*
Table-Driven Test
You want to test your code with many inputs and
expected result of these inputs.
The best approach is creating an array for inputs
and run tests each item in array and get expected result.

*/

func Test_SayHello_ValidArgument(t *testing.T) {

	inputs := []struct {
		name   string
		result string
	}{
		{name: "Dinesh", result: "Hello Dinesh"},
		{name: "Pushan", result: "Hello Pushan"},
		{name: "Yemek", result: "Hello Yemek!"},
	}

	for _, item := range inputs {
		result := sayHello(item.name)
		if result != item.result {
			t.Errorf("\"sayHello('%s')\" failed, expected -> %v, got -> %v", item.name, item.result, result)
		} else {
			t.Logf("\"sayHello('%s')\" succeded, expected -> %v, got -> %v", item.name, item.result, result)
		}
	}

}
