package h

import "fmt"

func PrintHello(name string) string {
	h := "Hello " + name
	fmt.Printf("\nGreeting is - %s\n", h)
	return h
} 

