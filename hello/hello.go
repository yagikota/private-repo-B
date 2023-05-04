package hello

import "fmt"

func hello1(s string) string {
	return fmt.Sprintf("func hello1: arg=%s", s)
}

func hello2(s string) string {
	return fmt.Sprintf("func hello2: arg=%s", s)
}
