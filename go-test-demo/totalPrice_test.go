package go_test_demo

import "testing"

func TestTotalPrice(t *testing.T) {
	// test case 1

	expected := uint(0)
	actual := TotalPrice(0, 150, 12)
	if expected != actual {
		t.Errorf("Expected %d doesnot match actual %d", expected, actual)
	} else {
		t.Logf("Succeded, expected -> %v, got -> %v", expected, actual)
	}

	// test case 2
	expected = uint(113)
	actual = TotalPrice(1, 100, 12)
	if expected != actual {
		t.Errorf("Expected %d does not match actual %d", expected, actual)
	} else {
		t.Logf("Succeded, expected -> %v, got -> %v", expected, actual)
	}

	// test case 3
	expected = uint(224)
	actual = TotalPrice(2, 100, 12)
	if expected != actual {
		t.Errorf("Expected %d does not match actual %d", expected, actual)
	} else {
		t.Logf("Succeded, expected -> %v, got -> %v", expected, actual)
	}
}

func Test_totalPrice(t *testing.T) {
	type parameters struct {
		nights  uint
		rate    uint
		cityTax uint
	}
	type testCase struct {
		name string
		args parameters
		want uint
	}
	tests := []testCase{
		{
			name: "test 0 nights",
			args: parameters{nights: 0, rate: 150, cityTax: 12},
			want: 0,
		},
		{
			name: "test 1 nights",
			args: parameters{nights: 1, rate: 100, cityTax: 12},
			want: 113,
		},
		{
			name: "test 2 nights",
			args: parameters{nights: 2, rate: 100, cityTax: 12},
			want: 224,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TotalPrice(tt.args.nights, tt.args.rate, tt.args.cityTax); got != tt.want {
				t.Errorf("totalPrice() = %v, want %v", got, tt.want)
			}
		})
	}
}
