package util

import "fmt"

func GetString(name string) string {
	h := "Hello " + name
	fmt.Printf("New Greeting is - %s\n", h)
	return h
}
