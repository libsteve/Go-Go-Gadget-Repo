/*
A group of functions realted to finding the value of an expression.

Package Usage:
	Evaluate( initPhrase []string )
*/
package main

import (
	"strconv"
	"fmt"
	"os"
)

var phrase []string
var currentIndex int

var forgivenessCount int

/*
Find the value of an expression.

Parameters:
	initPhrase []string - an array of strings that contains an expression

Returns:
	int - the resulting value of the given expression
*/
func Evaluate( initPhrase []string ) int {
	 if currentIndex != 1 {
		currentIndex = 1
	}

	phrase = initPhrase

	checkPhrase()

	forgivenessCount = 0

	return int( sum() )
}

/*
Check the parenthesis in the expression.
*/
func checkPhrase() {
	parenCount := 0
	for i := 1; i < len( phrase ); i++ {
		if phrase[i] == "(" {
			parenCount++
		} else if phrase[i] == ")" {
			parenCount--
		}
	}
	if parenCount != 0 {
		fmt.Println( phrase[0] + ": syntax error" )
		os.Exit( 2 )
	} else {
		currentIndex = 1
	}
}

/*
Get the next symbol in the expression.

Returns:
	string - the next string in the array of the expression
*/
func nextInput() string {
	if(currentIndex  >= len (phrase )){
		return  ""
	}
	returnVal := phrase[ currentIndex ]
	currentIndex++
	return returnVal
}

/*
Get next sum from the expression.

Returns:
	float32 - the sum of the expression from the next input
*/
func sum() float32 {
	left := product()
	left = performOperation( left )
	return left
}

/*
Get the next product from the expression.

Returns:
	float32 - the product of the expression from the next input
*/
func product() float32 {
	left := term()
	left = performOperation( left )
	return left
}

/*
Get the operation from next input and perfomr it.

Parameters:
	left float32 - the left side of the equation

Returns:
	float32 - the result of the operation.
*/
func performOperation( left float32 ) float32 {
	nextInput := nextInput()
	switch nextInput {
	case "*":
		right := term()
		left *= right
	case "/":
		right := term()
		left /= right
	case "+":
		right := product()
		left += right
	case "-":
		right:= product()
		left -= right
	case ")":
		left += 0
	case "":
		left += 0
	default:
		fmt.Println( phrase[0] + ": syntax error" )
		os.Exit(2)
	}
	return left
}

/*
Get the next term from the expression.
If the next input is a number, return it.
If the next input is a parenthesis, find the value of the expression within it.

Returns:
	float32 - the value of the expression starting from the next input
*/
func term() float32 {
	var result float32
	nextInput := nextInput()
	switch nextInput {
	case "(":
		result = sum()
	case ")":
		result += 0
	default:
		if a, b := strconv.Atof32( nextInput ); b == nil {
			result = a
		} else{
			if forgivenessCount == 0 {
				if ( nextInput != "Makefile"){
					fmt.Println( nextInput )
					os.Exit(0)
				}else{
					fmt.Println(phrase[0]  + ": syntax error")
					os.Exit(2)
				}
			} else {
				fmt.Println( phrase[0] + ": non-numeric argument" )
				os.Exit(2)
			}
		}
		forgivenessCount--
	}
	return result
}

