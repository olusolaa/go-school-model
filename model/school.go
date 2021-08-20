package model


type School struct {
	Name string
	Principal
	Teachers []Teacher
	Students []Student
	Staff [] Staff
}

func CalculateGPA(student Student) int {
	sum := 0
	for _, course := range student.Courses {
		sum += course.Score
	}
	return sum/len(student.Courses)
}





