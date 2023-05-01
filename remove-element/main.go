package main

import "fmt"

func removeLeftRight(arr []any, position string) []any {
	//yout code here!
	switch {
		case position == "left":
			return arr[1:]

		case position == "right":
			return arr[:len(arr)-1]

		default :
			return nil
	}
}

func main() {
	fmt.Println(removeLeftRight([]any{1, 2, 3, 4, 5}, "left")...)
	fmt.Println(removeLeftRight([]any{1, 2, 3, 4, 5}, "right")...)
	fmt.Println(removeLeftRight([]any{"Edo", "Budi", "Joko", "Tono"}, "left")...)
	fmt.Println(removeLeftRight([]any{"Edo", "Budi", "Joko", "Tono"}, "right")...)

}
