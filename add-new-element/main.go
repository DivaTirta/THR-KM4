package main

import "fmt"

func AddElement(data []int, newData int, position string) []int {
	switch {
		case position == "up" :
			return append([]int{newData}, data...)

		case position == "down":
			return append(data, newData)

		default:
			return nil 
	}
}

func main() {
	fmt.Println(AddElement([]int{1,2,3,4,5}, 6, "up"))
	fmt.Println(AddElement([]int{1,2,3,4,5}, 6, "down"))
}
