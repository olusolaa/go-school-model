package main

import (
	"fmt"
	"github.com/olusolaa/model"
)

func main() {
	applicant1 := model.Applicant{Name: "Ayo", Age: 2, ClassName: "jss1"}
	applicant2 := model.Applicant{Name: "James",Age: 13, ClassName: "jss1"}
	applicant3 := model.Applicant{Name: "Goodluck" , Age:15, ClassName: "jss3"}
	principal := model.Principal{Staff: model.Staff{Name: "Toluwase", Role:   "principal"},
		SchoolName: "BestWay"}

	school2 := model.School{}
	school := model.School{Name: "Best Way", Principal: principal}
	school.Admit(applicant1, &school.Students)
	school.Admit(applicant2, &school.Students)
	fmt.Println(school.Students, len(school.Students))
	school2.Admit(applicant3, &school2.Students)
	fmt.Println(school2.Students, len(school2.Students))
	school.Admit(applicant3, &school.Students)
	fmt.Println(school.Students, len(school.Students))

	student1 := model.Student{Applicant: applicant1, ID: 4, Courses: []model.Course{}}
	student2 := model.Student{Applicant: applicant3, ID: 4, Courses: []model.Course{}}
	course1 := model.Course{Name:"Biology", Class: "jss1"}
	student1.TakeCourse(course1)
	student2.TakeCourse(course1)
	fmt.Println(student2)
	fmt.Println(student1)
	student1.TakeCourse(course1)
	fmt.Println(student1)
}

