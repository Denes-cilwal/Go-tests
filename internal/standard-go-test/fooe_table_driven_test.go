package main

import "testing"

// table driven test
// Table tests reuse the same assertion logic, keeping your test DRY.
// TestFooerTableDriven function runs four subtests, one for each row of the table.
// The execution loop calls t.Run(), which defines a subtest.
// As a result, each row of the table defines a subtest named [NameOfTheFuction]/[NameOfTheSubTest].
func TestFooerTableDriven(t *testing.T) {
	// Defining the columns of the table
	var tests = []struct {
		name  string
		input int
		want  string
	}{
		// add test case here - subtest
		// the table itself
		{"9 should be Foo", 9, "Foo"},
		{"3 should be Foo", 3, "Foo"},
		{"1 is not Foo", 1, "1"},
		{"0 should be Foo", 0, "Foo"},
	}
	// The execution loop
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := Fooer(tt.input)
			// assertion logic
			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}
