package funcs

import "fmt"

func InputID() {
	var nameID int16
	var sec byte
	var name string
	fmt.Print("Введите имя сотрудника :")
	fmt.Scan(&name)
	fmt.Println("")
	fmt.Print("Введите новый ID состояйщий из 4 цифр:")
	_, err := fmt.Scan(&nameID)
	if err != nil {
		return
	}
	fmt.Println("")
	fmt.Print("Введите уровень доступа от 1-3:")
	_, err = fmt.Scan(&sec)
	if err != nil {
		return
	}
	if sec == 1 {
		nameID += 10000

	} else if sec == 2 {
		nameID += 20000

	} else if sec == 3 {
		nameID += 30000

	} else {
		fmt.Println("Не правильный код доступа")
	}
	ShowAll()
}
