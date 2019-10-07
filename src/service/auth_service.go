package service

import "fmt"

func hello(term string) string  {

	if len(term) == 0 {
		return "Hello Done"
	}

	return fmt.Sprintf("Hello %v", term)
}