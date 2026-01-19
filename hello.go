package main

import "fmt"

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
	fmt.Print(nums, nums2, rgb)
}

// variable
