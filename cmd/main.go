package main

import (
	"fmt"

	"job4j.ru/go-lang-base/internal/base"
)

func main() {
	first := 100
	second := 10
	res := base.Max(first, second)
	fmt.Println(fmt.Sprintf("%d + %d = %d", first, second, res))
}
