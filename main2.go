package main

import "fmt"

func main() {
	//1 задание
	var temperature int = 15

	if temperature < 0 {
		fmt.Println("Холодно")
	} else if temperature >= 0 && temperature <= 20 {
		fmt.Println("Тепло")
	} else {
		fmt.Println("Жарко")
	}
	//2 задание
	var score int = 85

	if score >= 90 {
		fmt.Println("Отлично")
	} else if score >= 70 && score <= 89 {
		fmt.Println("Хорошо")
	} else if score >= 50 && score <= 69 {
		fmt.Println("Удовлетворительно")
	} else {
		fmt.Println("Не сдал")
	}
	//3 задание
	var hour int = 14

	switch {
	case hour >= 0 && hour <= 5:
		fmt.Println("Ночь")
	case hour >= 6 && hour <= 11:
		fmt.Println("Утро")
	case hour >= 12 && hour <= 17:
		fmt.Println("День")
	case hour >= 18 && hour <= 23:
		fmt.Println("Вечер")
	default:
		fmt.Println("Некорректное значение часа")
	}
	//4 задание
	var number int = 7

	if number%2 == 0 {
		fmt.Println("Чётное число")
	} else {
		fmt.Println("Нечётное число")
	}
	//5 задание
	var day string = "Monday"

	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("Будний день")
	case "Saturday", "Sunday":
		fmt.Println("Выходной")
	default:
		fmt.Println("Некорректный день")
	}
	//6 задание
	var balance int = 1500

	if balance >= 0 {
		fmt.Println("Баланс положительный")
	} else {
		fmt.Println("Баланс отрицательный")
	}
	//7 задание
	var age int = 20

	if age < 13 {
		fmt.Println("Ребёнок")
	} else if age >= 13 && age <= 17 {
		fmt.Println("Подросток")
	} else {
		fmt.Println("Взрослый")
	}
	//8задание
	var command string = "start"

	switch command {
	case "start":
		fmt.Println("Система запускается...")
	case "stop":
		fmt.Println("Система останавливается...")
	case "restart":
		fmt.Println("Система перезапускается...")
	default:
		fmt.Println("Неизвестная команда")
	}
	//9 задание
	var grade int = 4

	switch grade {
	case 5:
		fmt.Println("A")
	case 4:
		fmt.Println("B")
	case 3:
		fmt.Println("C")
	case 2:
		fmt.Println("D")
	case 1:
		fmt.Println("F")
	default:
		fmt.Println("Некорректная оценка")
	}
}

//
