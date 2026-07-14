package main

import "fmt"

func main() {

	nums := []int{20, 30, 40, 50, 60, 70, 80, 90, 100}

	var index int = 2

	nums = append(nums[:index], nums[index+1:]...)

	fmt.Println(nums)

}
