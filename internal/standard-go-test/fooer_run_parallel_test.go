package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

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

/* B is a type passed to Benchmark functions to manage benchmark */

// timing and to specify the number of iterations to run.
func BenchmarkFooer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fooer(i)

	}
}

/*
Testify provides assert functions and mocks, which are similar to traditional testing frameworks, like JUnit for Java or Jasmine for Node.js.
Testify provides two packages, require and assert. The require package will stop execution if there is a test failure, which helps you fail fast.
assert lets you collect information, but accumulate the results of assertions
*/
func TestFooerWithTestify(t *testing.T) {
	// assert equality
	assert.Equal(t, "Foo", Fooer(0), "0 is divisible by 3, should return Foo")
	// assert inequality
	assert.NotEqual(t, "Foo", Fooer(1), "1 is not divisible by 3, should not return Foo")
}

func TestMapWithTestify(t *testing.T) {

	// require equality
	require.Equal(t, map[int]string{1: "1", 2: "2"}, map[int]string{1: "1", 2: "3"})

	// assert equality
	assert.Equal(t, map[int]string{1: "1", 2: "2"}, map[int]string{1: "1", 2: "2"})
}

/*
The output log also clearly indicates the difference between the actual output and the expected output.
Compared to Go’s built-in testing package, the output is more readable, especially when the testing data is complicated,
such as with a long map or a complicated object. The log points out exactly which line is different, which can boost your productivity.
*/
