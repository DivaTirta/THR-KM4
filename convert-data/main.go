package main

import (
	"fmt"
	"strconv"
	"strings"
)

type User struct {
	Name    string
	Age     int
	Address string
}

func ConvertData(data string) User {
	//your code here!
	separate := strings.Split(data, ",")
	age, _ := strconv.Atoi(separate[1])
	return User{Name: "name: " + separate[0] + ", age:" , Age: age , Address: ", address: " +  separate[2]}
}

func main() {
	fmt.Println(ConvertData("Edo,20,Jakarta"))
	fmt.Println(ConvertData("Budi,30,Bandung"))

}
