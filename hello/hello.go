package hello

import "fmt"

func Hello1(s string) string {
	return fmt.Sprintf("func hello1: arg=%s", s)
}

func Hello2(s string) string {
	return fmt.Sprintf("func hello2: arg=%s", s)
}
