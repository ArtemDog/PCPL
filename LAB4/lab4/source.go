package main

import "fmt"

type Student struct {
	Id         int
	FirstName  string
	SecondName string
	Age        int
	Faculty    string
	Course     int
}

func NewStudent(id int, firstName, secondName string, age int, faculty string, course int) Student {
	return Student{
		Id:         id,
		FirstName:  firstName,
		SecondName: secondName,
		Age:        age,
		Faculty:    faculty,
		Course:     course,
	}
}

func AddStudent(students *[]Student, student Student) {
	*students = append(*students, student)
}

func RemoveStudent(students *[]Student, id int) {
	for i, student := range *students {
		if student.Id == id {
			*students = append((*students)[:i], (*students)[i+1:]...)
			return
		}
	}
}

func UpdateStudent(students *[]Student, id int, updatedStudent Student) {
	for i, student := range *students {
		if student.Id == id {
			(*students)[i] = updatedStudent
			return
		}
	}
}

func PrintAllStudents(students []Student) {
	for i := 0; i < len(students); i++ {
		fmt.Printf("Id: %d\nFirst Name: %s\nSecond Name: %s\nAge: %d\nFaculty: %s\nCourse: %d\n", students[i].Id, students[i].FirstName, students[i].SecondName, students[i].Age, students[i].Faculty, students[i].Course)
		fmt.Println("--------------")
	}
}

func SortStudentsByCourse(students []Student) []Student {
	var Sorted []Student = students
	for i := 0; i < len(Sorted)-1; i++ {
		for j := i + 1; j < len(Sorted); j++ {
			if Sorted[i].Course > Sorted[j].Course {
				Sorted[i], Sorted[j] = Sorted[j], Sorted[i]
			}
		}
	}
	return Sorted
}
