package hometaskes

import "fmt"

func Split() {
	fmt.Println("Вы выбрали задание разбиение кода на несколько функций. \n" +
		"Теперь данная программа вычисляет и площадь и периметр прямоугольника\n" +
		"Введите длину и ширину прямоугольника")
	var RectangleLengh, RectangleWidth int
	fmt.Scan(&RectangleWidth, &RectangleLengh)

	Perimertr(RectangleLengh, RectangleWidth)
	Area(RectangleLengh, RectangleWidth)

}
func Perimertr(length, width int) {
	fmt.Println("Периметр прямоугольника равен:", 2*(length+width))
}
func Area(length, width int) {
	fmt.Println("Площадь прямоугольника равна", width*length)
}
