package hometaskes

import (
	"fmt"
	"log"
)

func Square() {
	fmt.Println("Вы выбрали задание вычисление периметра квадрата\n" +
		"Введите длину стороны")

	var A int
	fmt.Scan(&A)
	fmt.Println("Периметр квадрата равен", 4*A)
}

func RectanglePerimetr() {
	fmt.Println("Вы выбрали задание вычисление периметра прямоугольника\n" +
		"Введите длину")
	var RectangleLengh, RectangleWidth int
	fmt.Scan(&RectangleLengh)
	fmt.Println("Введите ширину")

	fmt.Scan(&RectangleWidth)
	fmt.Println("Периметр прямоугольника равна", 2*(RectangleLengh+RectangleWidth))

}
func Jubilee() {
	var Age int
	log.Println("Программа запущена")
	fmt.Println("Вы выбрали задание Юбилей. Введите свой возраст:")
	fmt.Scan(&Age)
	const Five int = 5
	fmt.Println("Вам осталолось", Five-Age%Five, "лет до вашего юбилея.")
	log.Println("Программа завершила свою работу.")
}
func RectangleArea() {
	var RectangleLengh, RectangleWidth int
	fmt.Println("Вы выбрали задание вычисление площади прямоугольника\n" +
		"Введите длину")
	fmt.Scan(&RectangleLengh)
	fmt.Println("Введите ширину")
	fmt.Scan(&RectangleWidth)
	fmt.Println("Площадь прямоугольника равна", RectangleWidth*RectangleLengh)

}
