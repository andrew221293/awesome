package Sum

import "fmt"

func Sum(a, b int) error{
	result := a + b
	if result < 10 {
		return fmt.Errorf("is minor")
	}
	return nil
}