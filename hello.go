package hutil

import "fmt"
import "github.com/mohan-maturi/hutil/util"
import "github.com/mohan-maturi/wutil/w"

func PrintHello(name string) string {
	return util.GetString(name)
} 

func PrintWorld(name string) string {
	fmt.Printf("Calling world greeting for %s\n", name)
	return w.PrintWorld(name)
} 
