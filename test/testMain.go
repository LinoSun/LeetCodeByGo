package main

import (
	"fmt"
	"goleetcode/stack"
)

func main() {
	nums := []int{-2,1,2,-2,1,2}
	fmt.Println(stack.Find132pattern(nums))
}
