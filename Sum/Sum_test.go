package Sum

import (
	"testing"
)

func TestSum(t *testing.T) {
	tables := []struct{
		a, b int
		expectedError bool
	}{
		{1, 2, true},
		{10, 5, false},
	}

	for _, item := range tables{
		err := Sum(item.a, item.b)
		if item.expectedError == false && err != nil {
			t.Errorf("No error should occur")
		}
		if item.expectedError == true && err == nil {
			t.Errorf("An error was expected")
		}
	}
}