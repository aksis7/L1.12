package main

import (
	"fmt"
)

func main() {
	// Исходная последовательность строк
	strings := []string{"cat", "cat", "dog", "cat", "tree"}

	// Создаем множество
	set := make(map[string]struct{})

	// Добавляем элементы в множество
	for _, s := range strings {
		set[s] = struct{}{}

	}

	// Вывод множества
	fmt.Println("Множество:")
	for key := range set {
		fmt.Println(key)
	}
}
