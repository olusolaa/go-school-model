package model

import "fmt"

type Student struct{
	Applicant
	ID int
	Courses []Course
}

func (student *Student) TakeCourse(course Course)  {
	if student.ClassName == course.Class{
		student.Courses = append(student.Courses, course)
	}else {
		fmt.Printf("%s, cannot take this course, it is meant for %s students", student.Name, course.Class)
	}
}


