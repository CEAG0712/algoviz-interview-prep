package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestAddTwoNumbers(t * testing.T) {
	result := Add(2,3)
	assert.Equal(t,5,result,"Adding 2 + 3 should equal 5" )
}

func TestSubstractTwoNumbers(t *testing.T){
	result := Substract(5,3)
	assert.Equal(t, 2, result, "Substracting 5 - 3 should equal 2")
}

func TestDivideByZero(t *testing.T){
	result, err := Divide(10, 0)
	assert.Error(t, err,"Dividing by zero should return an error" )
	assert.Equal(t, 0, result, "Result should be 0 when there's an error")
}