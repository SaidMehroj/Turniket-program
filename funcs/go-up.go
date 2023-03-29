package funcs

import "fmt"

func ViewFloor() {
	fmt.Print("Введите ваш ID: ")
	var sign int64
	fmt.Scan(&sign)
	if sign/10000 == 1 {
		fmt.Println("Вы можете поднятся только на первый этаж")
		ShowAll()
	} else if sign/10000 == 2 {
		fmt.Println("Вы можете поднятся на первый и второй этаж")
		ShowAll()
	} else if sign/10000 == 3 {
		fmt.Println("Вы можете поднятся на все этажи")
		ShowAll()
	} else {
		fmt.Println("Неправильный ID")
		ShowAll()
	}
}
