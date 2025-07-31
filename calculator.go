package main

import "fmt"

func Add(a int, b int) int {
    return a + b
}

func Substract(a int, b int)int{
	return a - b
}

func Divide(a int, b int) (int, error){
	if b==0{
		return 0, fmt.Errorf("cannot divide by zero")
	}

	return a/b, nil 
}