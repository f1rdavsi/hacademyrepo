package univer

import (
	"fmt"
	"math/rand"
	university "university/exams"

	"university/student"
)

func Univer() {
	var Student student.Candidate = student.Candidate{
		Name:  "Мистер Бист",
		Age:   rand.Intn(26) + 18, // Студент может быть в возрасте от 18 до 43 лет
		Score: 0,
		Subjects: student.Subjects{
			Math:       0,
			Physics:    0,
			Chemistry:  0,
			Literature: 0,
		},
	}

	fmt.Println("Студент", Student.Name, "пришёл на экзамен.")
	fmt.Println("Студенту", Student.Name, "возраст", Student.Age, "года.")

	fmt.Println("Студент", Student.Name, "начинает сдавать экзамены по математике, физике, химии и литературе.")

	Student = university.TakeExams(Student)

	fmt.Println("Студент", Student.Name, "закончил сдачу экзаменов.")
	fmt.Println("Студент", Student.Name, "получил общую оценку:", Student.Score)

	fmt.Println("Студент", Student.Name, "проверяет результаты экзаменов.")

	university.CheckResults(Student)
}
