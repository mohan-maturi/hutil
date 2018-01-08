package h

import "fmt"
import "github.com/mohan-maturi/wutil/w"

func PrintHello(name string) string {
	h := "Hello " + name
	fmt.Printf("New Greeting is - %s\n", h)
	return h
} 

func PrintWorld(name string) string {
	fmt.Printf("Calling world greeting for %s\n", name)
	return w.PrintWorld(name)
} 
