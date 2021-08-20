package model

import "fmt"

type Teacher struct {
	Staff
	Course Course
}

func (teacher Teacher) GradeStudent(student *Student, score int)  {
	for i, course := range student.Courses {
		if teacher.Course.Name == course.Name && teacher.Course.Class == course.Class{
			student.Courses[i].Score = score
			return
		}
	}
	fmt.Printf("student doe not offer %v \n", teacher.Course)
}