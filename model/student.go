package model

type Student struct{
	Applicant
	ID int
	Courses []Course
}

func (student *Student) TakeCourse(course Course)  {
	if student.ClassName == course.Class{
		student.Courses = append(student.Courses, course)
	}
}

