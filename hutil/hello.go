package hutil

import "fmt"

func PrintHello(name string) string {
	h := "Hello " + name
	fmt.Printf("Greeting is - %s", h)
	return h
} 

