package main

import (
	"fmt"
	// "first"
)


func main() {
	var planets [8]string

	fmt.Println(planets)
	fmt.Println("Привет, Мир!")

	nums := []int{10, 20, 30, 40, 50}
	nums2 := nums

	nums[1] = 5

	rgb := [3][3]int{
		{0, 0, 0},
		{128, 128, 128},
		{255, 255, 255},
	}

	name := "Alex"

	fmt.Print(nums, nums2, rgb)
	fmt.Printf("My name %s", name)
	// first.MyPrint("Hello, World!")

	// myFunc  // camel case
	// MyFunc  // Pascal case
	// my_func // snake case
	// my-func  // kebab case
}
