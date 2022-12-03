package main

import (
	"fmt"
)

type member struct {
	name string
	age  int
}

func main() {
	var yuYuan = member{
		name: "Yu Yuan",
		age:  24,
	}

	fmt.Println("歡迎問候語：", GetGreeting(yuYuan.name))

	fmt.Println("開始檢查年齡是否成年...")
	if IsAdult(yuYuan.age) {
		fmt.Println(yuYuan.name, yuYuan.age, "歲成年。")
	} else {
		fmt.Println(yuYuan.name, yuYuan.age, "歲未成年。")
	}
}

func GetGreeting(name string) string {
	return fmt.Sprintf("Hello, %s! Nice to meet you!", name)
}

func IsAdult(age int) bool {
	return age >= 18
}
