package main

import "testing"

func Test_sum_of_two_digits(t *testing.T) {

	expectedValue := 6
	actualValue := sum(2, 4)

	if expectedValue != actualValue {
		t.Errorf("expected value %d is not equals to actual value %d\n", expectedValue, actualValue)
	}
}

// this will be production code
func sum(a, b int) int {
	return a + b
}
