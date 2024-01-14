package fuzz

/*
Fuzzing can be used anywhere you have input from an untrusted source, whether that’s on the open internet or somewhere security sensitive.
Fuzzing can be used to test software, files, network policies, applications, and libraries.
For example, say you have an application that requires the user’s name as input, and you have to test the application for invalid data.
You can create a test case with fuzzing that uses input with special characters or values that can lead to crashes or result in a memory or buffer overflow.
*/

// fuzz testing in modules

func Equal(a []byte, b []byte) bool {
	// Check if the slices are of the same length
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		// Check if the elements in the same index are the same
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
