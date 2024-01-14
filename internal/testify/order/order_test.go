package order

import (
	"github.com/Rhymond/go-money"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestComputeTotal(t *testing.T) {
	// running multiple test
	// running two subset in parallel
	t.Run("nominal case", func(t *testing.T) {
		t.Parallel()
		o := Order{
			ID:                "45",
			CurrencyAlphaCode: "EUR",
			Items: []Item{
				{
					ID:        "458",
					Quantity:  2,
					UnitPrice: money.New(200, "EUR"),
				},
			},
		}

		total, err := o.ComputeTotal()
		assert.NoError(t, err)
		assert.Equal(t, int64(200), total.Amount())
		assert.Equal(t, "EUR", total.Currency().Code)

	})

	t.Run("currency issues", func(t *testing.T) {
		t.Parallel()
		o := Order{
			ID:                "46",
			CurrencyAlphaCode: "USD",
			Items: []Item{
				{
					ID:        "459",
					Quantity:  2,
					UnitPrice: money.New(200, "EUR"),
				},
			},
		}

		// no need to check the value of total
		_, err := o.ComputeTotal()
		assert.NoError(t, err)

	})

	/*
		Testify
			- easy assertions -
			- Mocking
			- Testing suite interfaces and functions
	*/

	/*
		The assert.Equal function needs three main things to do its job:
			- The Toolbox (t): This is *testing.T. It's like telling the assertion, "Here's a toolbox you can use if you need to report a problem or do something with the test."
			- Expected Value (200): This is what you expect the result should be.
			- Actual Value (total.Amount()): This is the actual result produced by your code.
			- When assert.Equal runs, it compares the actual value with the expected value.
			- If they don't match, it uses the toolbox (t) to report a problem. If you didn't give it this toolbox,
			- it wouldn't know how to tell you that something went wrong in the test.
	*/

	/*
		Test Control: It provides control over the test execution. For example, if an assertion fails, you can use methods on the *testing.T object to fail the test immediately (using t.FailNow()), or continue with further checks.
		Contextual Information: The *testing.T object carries contextual information about the test, such as its name, whether it's part of a larger suite, and other state information.
		This can be useful for more complex test setups.
	*/

}
