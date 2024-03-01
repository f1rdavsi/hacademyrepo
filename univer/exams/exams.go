package university

import (
	"fmt"
	"math/rand"

	"university/student"
)

func TakeExams(candidate student.Candidate) student.Candidate {

	// Проверка возраста студента
	if candidate.Age < 18 {
		fmt.Println("Студент", candidate.Name, "слишком молод для сдачи экзамена!")
		return student.Candidate{}
	} else if candidate.Age > 43 {
		fmt.Println("Студент", candidate.Name, "слишком стар для сдачи экзамена!")
		return student.Candidate{}
	}

	// Студент сдаёт экзамены по разным предметам
	candidate.Subjects.Math = rand.Intn(101)
	candidate.Subjects.Physics = rand.Intn(101)
	candidate.Subjects.Chemistry = rand.Intn(101)
	candidate.Subjects.Literature = rand.Intn(101)

	// Общая оценка вычисляется как среднее арифметическое
	candidate.Score = (candidate.Subjects.Math + candidate.Subjects.Physics + candidate.Subjects.Chemistry + candidate.Subjects.Literature) / 4

	fmt.Println("Студент", candidate.Name, "сдал экзамены успешно!")

	return candidate
}

func CheckResults(candidate student.Candidate) {
	// Проверка результатов студента
	if candidate.Score < 60 {
		fmt.Println("Студент", candidate.Name, "не сдал экзамен! Оценка:", candidate.Score)
	} else {
		fmt.Println("Студент", candidate.Name, "сдал экзамен успешно! Общая оценка:", candidate.Score)
	}
}
