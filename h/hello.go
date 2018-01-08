package h

import "fmt"

func PrintHello(name string) string {
	h := "Hello " + name
	fmt.Printf("Greeting is - %s\n", h)
	return h
} 

