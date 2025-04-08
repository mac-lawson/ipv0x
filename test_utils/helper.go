package testutils

import (
	"reflect"
	"testing"
)

// Assert fails the test if the condition is false.
// Usage:
//
//	testhelper.Assert(t, condition, "error message")
//
// Parameters:
//
//	tb:        Testing interface (usually *testing.T).
//	condition: Boolean condition to evaluate.
//	msg:       Message to display if the assertion fails.
func Assert(tb testing.TB, condition bool, msg string) {
	tb.Helper()
	if !condition {
		tb.Fatalf("Assertion failed: %s", msg)
	}
}

// Ok fails the test if an err is not nil.
// Usage:
//
//	err := someFunction()
//	testhelper.Ok(t, err)
//
// Parameters:
//
//	tb:  Testing interface (usually *testing.T).
//	err: Error to check.
func Ok(tb testing.TB, err error) {
	tb.Helper()
	if err != nil {
		tb.Fatalf("Unexpected error: %s", err)
	}
}

// Equals fails the test if exp is not equal to act.
// Usage:
//
//	testhelper.Equals(t, expectedValue, actualValue)
//
// Parameters:
//
//	tb:  Testing interface (usually *testing.T).
//	exp: Expected value.
//	act: Actual value.
func Equals(tb testing.TB, exp, act interface{}) {
	tb.Helper()
	if !reflect.DeepEqual(exp, act) {
		tb.Fatalf("Expected %v but got %v", exp, act)
	}
}
