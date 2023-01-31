package go_test_demo

import "fmt"

func sayHello(name string) string {
	// empty
	if len(name) == 0 {
		return "Hello Anonymous!"
	}

	return fmt.Sprintf("Hello %s", name)
}
