package main

import "fmt"

func main() {

	s := []int{1, 2}

	s = append(s, 0, 7, 3, 4, 5)

	fmt.Println(s)

}
