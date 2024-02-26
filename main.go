package main

import (
	"fmt"
	"log"

	"main.go/hometaskes"
	"main.go/university"
)

func main() {
	log.Println("Программа начала работу")
	var day, task int
	fmt.Println("Выберите день:")
	fmt.Scan(&day)
	switch day {
	case 1: // Дз на первый день
		fmt.Println("День знакомства. На первый день не было дз :)")

	case 2: // дз на второй день
		{
			fmt.Println("Вы выбрали второй день. Выберите задачу:\n" +
				"1.Периметр квадрата.\n" +
				"2.Периметр прямоугольника\n" +
				"3.Площадь прямоугольника\n" +
				"4.Юбилей")
			fmt.Scan(&task)
			switch task {
			case 1: // Периметр квадрата
				hometaskes.Square()
			case 2: //Периметр прямоугольника
				hometaskes.RectanglePerimetr()
			case 3: // Площадь прямоугольника
				hometaskes.RectangleArea()
			case 4: // Юбилей
				hometaskes.Jubilee()
			default:
				fmt.Println(" Тагоко задания на второй день не было")
			}
		}
	case 3: //Дз на третий день
		{
			fmt.Println("Вы выбрали третий день. Выберите задачу:\n" +
				"1.Разбиение кода на разные пакеты.")
			fmt.Scan(&task)
			if task == 1 {
				hometaskes.Split()
			} else {
				fmt.Println("Такого задания на этот день не было")
			}
		}
	case 4:
		{ //Дз на четвертый день
			fmt.Println("Вы выбрали четвертый день.Выберите задачу:\n" +
				"1.Собрание всех домашних заданий в одну программу\n" +
				"2.Вывод таблицы умножения\n" +
				"3.Числа Фибоначчи")
			fmt.Scan(&task)
			if task == 1 { //Собрание всех домашних заданий в одну программу
				fmt.Println("Об этой же программе и идет речь)")
			} else if task == 2 { // Вывод таблицы умножения
				hometaskes.Multiplication()
			} else if task == 3 {
				hometaskes.Fibonachi()

			} else {
				fmt.Println("Такого дз на этот день не было!")
			}
		}
	case 5:
		{
			fmt.Println("Вы выбрали пятый день. Выберите задачу:\n" +
				"Универститет")

			university.Univer()
		}
	case 6:
		{

		}
	case 7:
		{

		}

	default:
		fmt.Println("Такого дз еще не было!")

	}
	log.Println("Программа завершена")
}
