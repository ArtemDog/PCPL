package main

import "fmt"

func main() {
	var students []Student

	student1 := NewStudent(1, "Артем", "Кузнецов", 19, "ИУ5", 3)
	student2 := NewStudent(2, "Алексей", "Князев", 19, "ИУ5", 2)
	student3 := NewStudent(3, "Дмитрий", "Егоров", 19, "ИУ5", 2)

	AddStudent(&students, student1)
	AddStudent(&students, student2)
	AddStudent(&students, student3)
	fmt.Println("Все студенты:")
	PrintAllStudents(students)
	fmt.Printf("\n============================\n\n")

	fmt.Println("Обновим данные студента")
	updatedStudent := NewStudent(2, "Алексей", "Князев", 20, "ИУ5", 2)
	UpdateStudent(&students, 2, updatedStudent)
	PrintAllStudents(students)
	fmt.Printf("\n============================\n\n")

	fmt.Println("Удалим студента:")
	RemoveStudent(&students, 3)
	PrintAllStudents(students)
	fmt.Printf("\n============================\n\n")

	fmt.Println("Отсортируем студентов по курсу:")
	sortedStudents := SortStudentsByCourse(students)
	PrintAllStudents(sortedStudents)
}
