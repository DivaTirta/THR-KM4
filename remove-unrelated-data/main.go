package main

import "fmt"

func removeUnrelated(dataMap map[string]any, key string) map[string]any {
	//your code here
	delete(dataMap, key)
	return dataMap
}

func main() {
	fmt.Println(removeUnrelated(map[string]any{"name": "Edo", "age": 20, "address": "Jakarta"}, "address"))
	fmt.Println(removeUnrelated(map[string]any{"run": "lari", "jump": "loncat", "swim": "berenang"}, "jump"))
	fmt.Println(removeUnrelated(map[string]any{"satu": "ichi", "dua": "ni", "tiga": "san", "empat": "yon", "lima": "go"}, "tiga"))

}
