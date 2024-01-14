package main

import "testing"

/*
Running Parallel Tests
By default, tests are run sequentially;
the method Parallel() signals that a test should be run in parallel.
All tests calling this function will be executed in parallel. go test handles parallel tests by pausing each test that calls t.Parallel(),
and then resuming them in parallel when all non-parallel tests have been completed.
The GOMAXPROCS environment defines how many tests can run in parallel at one time, and by default this number is equal to the number of CPUs.



*/

func TestFooerParallel(t *testing.T) {
	t.Run("Test 3 in Parallel", func(t *testing.T) {
		t.Parallel()
		result := Fooer(3)
		if result != "Foo" {
			t.Errorf("Result was incorrect, got: %s, want: %s.", result, "Foo")
		}
	})
	t.Run("Test 7 in Parallel", func(t *testing.T) {
		t.Parallel()
		result := Fooer(7)
		if result != "7" {
			t.Errorf("Result was incorrect, got: %s, want: %s.", result, "7")
		}
	})
}
