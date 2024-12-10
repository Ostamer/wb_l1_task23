package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Println("Введите количество элементов в слайсе:")
	fmt.Scan(&n)

	slice := make([]int, n)
	fmt.Println("Введите элементы слайса:")
	for i := 0; i < n; i++ {
		fmt.Printf("Элемент %d: ", i+1)
		fmt.Scan(&slice[i])
	}

	var index int
	fmt.Println("Введите индекс для удаления (0 до", n-1, "):")
	fmt.Scan(&index)

	if index < 0 || index >= n {
		fmt.Println("Вы ввели индекс вне диапазона.")
		return
	}

	slice = append(slice[:index], slice[index+1:]...)

	fmt.Println("Слайс после удаления элемента:", slice)
}
