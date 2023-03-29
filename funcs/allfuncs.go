package funcs

import "fmt"

func ShowAll() {
	fmt.Println("Выберите действие:")
	fmt.Println("1 - Создать новый ID")
	fmt.Println("2 - Посмотреть пользователей")
	fmt.Println("3 - На какой этаж хотите поднятся")
	fmt.Println("4 - Выход из программы")
	var check byte
	fmt.Scan(&check)
	switch check {
	case 1:
		InputID()
	case 2:
		OutputID()
	case 3:
		ViewFloor()
	case 4:
		return
	}
}
