//remove element from slice

package main

import "fmt"

func main() {

	nums := []int{20, 30, 40, 50, 60, 70, 80, 90, 100}

	value := 40

	for i := range nums {
		if nums[i] == value {
			nums = append(nums[:i], nums[i+1:]...)
			break
		}
	}

	fmt.Println(nums)

}
