package main

import "fmt"

type house struct {
	city string
	ID   int
}

func newHouse(city string) house {

	h := house{
		city: city,
		ID:   0,
	}

	return h

}

func main() {

	houses := []string{}

	houses = append(houses, "name")

	fmt.Println(houses)

}
