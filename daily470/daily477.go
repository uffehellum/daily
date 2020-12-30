package main

import "fmt"

func main() {
	
	var functions [](func()int)
	for i := 0; i < 10; i++ {
		x := i
		functions = append(functions, func () int { return x})
	}
	for _, f := range(functions){
		fmt.Println(f())
	}
}