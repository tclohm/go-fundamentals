package utils

import "testing"

func TestAverage(t *testing.T) {
	expected := 4
	actual := utils.average(1, 2, 3)

	if actual != expected {
		t.Errorf("Average was incorrect! Expected: %d, Actual: %d", expected, actual)
	}
}