package main

import (
	"fmt"

	"github.com/gopherjs/gopherjs/js"
)

func main() {
	js.Module.Get("exports").Set("gcd", gcd)

}
func subUntillNegative(lhs int, rhs int) int {
	var left = lhs
	var right = rhs
	for left >= right {

		left -= right
	}
	return left
}
func prependInt(x []int, y int) []int {
	x = append(x, 0)
	copy(x[1:], x)
	x[0] = y
	return x
}
func gcd(p1 int, p2 int, verbose bool) int {
	var stack = make([]int, 4)
	/* Initializing the stack
	// Slice Shuold look like this

		B= Big Num
		S = Small num
		[S,B]
	*/
	if p1 > p2 {
		stack = prependInt(stack, p1)
		stack = prependInt(stack, p2)

	} else {
		stack = prependInt(stack, p2)
		stack = prependInt(stack, p1)
	}
	if verbose {
		fmt.Println("Initialized Stack", stack)
	}
	for {
		var residue = subUntillNegative(stack[1], stack[0])
		if verbose {
			fmt.Println("Res", residue)
		}
		if residue == 0 {
			if verbose {
				fmt.Println("Returned")
			}
			return stack[0]
		}
		if verbose {
			fmt.Println("Stack before push:", stack)
		}
		stack = prependInt(stack, residue)
		if verbose {
			fmt.Println("Stack after push:", stack)
		}

	}
}
