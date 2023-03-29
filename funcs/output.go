package funcs

import "fmt"

func OutputID() {
	fmt.Println("Выберите пользователя")
	fmt.Println("1 - John")
	fmt.Println("2 - Diana")
	fmt.Println("3 - Tony")
	fmt.Println("4 - Mary")
	fmt.Println("5 - Выход")
	var check byte
	fmt.Scan(&check)
	switch check {
	case 1:
		fmt.Printf("Вы имеете доступ ко всем этажам, ваш ID - %v", John)
		ShowAll()
	case 2:
		fmt.Printf("Вы имеете доступ к первому и второму этажу, ваш ID - %v", Diana)
		ShowAll()
	case 3:
		fmt.Printf("Вы имеете доступ к первому этажу, ваш ID - %v", Tony)
		ShowAll()
	case 4:
		fmt.Printf("Вы имеете доступ к первому этажу, ваш ID - %v", Mary)
		ShowAll()
	case 5:
		ShowAll()
	}
}
