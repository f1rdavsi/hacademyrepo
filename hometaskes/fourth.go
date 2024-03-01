package hometaskes

import "fmt"

func Multiplication() {
	fmt.Println("Вы выбрали задание Вывод таблицы умножения.")
	for i := 1; i < 10; i++ {
		fmt.Println("___________")
		for j := 1; j < 10; j++ {
			fmt.Println(i, "x", j, "=", i*j)
		}

	}

}
func Fibonachi() {
	var x, y int
	fmt.Println("Вы выбрали задание Вывод чисел Фибоначчи.\n" +
		"Введите начальное и конечное значение")
	fmt.Scan(&x, &y)
	fmt.Println("Числа Фибоначчи в диапазоне от", x, "до", y, ":")

	for i := 0; fibonacci(i) <= y; i++ {
		fib := fibonacci(i)
		if fib >= x {
			fmt.Println(fib)
		}
	}
}
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
