package main

import "testing"

/*
	Unit tests in Go are located in the same package (that is, the same folder) as the tested function.
	By convention, if your function is in the file fooer.go file, then the unit test for that function is in the file fooer_test.go.
*/

func TestFooer(t *testing.T) {
	result := Fooer(3)
	if result != "Foo" {
		t.Errorf("Result was incorrect, got: %s, want %s.", result, "Foo")
	}
}

func TestFooer2(t *testing.T) {
	input := 3
	result := Fooer(3)
	t.Logf("The input was %d", input)
	if result != "Foosß" {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, "Foo")
	}
	t.Fatalf("Stop the test now, we have seen enough")
	t.Error("This won't be executed")
}

func TestFooerSkiped(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	result := Fooer(3)
	if result != "Foo" {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, "Foo")
	}
}
