package main

import (
	"fmt"
	"github.com/olusolaa/model"
)

func main() {
	applicant1 := model.Applicant{Name: "Ayo", Age: 2, ClassName: "jss1"}
	applicant2 := model.Applicant{Name: "James",Age: 13, ClassName: "jss1"}
	applicant3 := model.Applicant{Name: "Goodluck" , Age:15, ClassName: "jss3"}
	staff := model.Staff{Name: "Toluwase", Role:   "principal"}
	principal := model.Principal{Staff: staff, SchoolName: "BestWay"}

	school2 := model.School{}
	school := model.School{Name: "Best Way", Principal: principal}
	school.Admit(applicant1, &school.Students)
	school.Admit(applicant2, &school.Students)
	fmt.Println(school.Students, len(school.Students))
	school2.Admit(applicant3, &school2.Students)
	fmt.Println(school2.Students, len(school2.Students))
	school.Admit(applicant3, &school.Students)
	fmt.Println(school.Students, len(school.Students))

	student1 := &school.Students[0]
	student2 := &school.Students[1]
	course1 := model.Course{Name:"Biology", Class: "jss1"}
	//course2 := model.Course{Name:"Biology", Class: "jss3"}
	course3 := model.Course{Name:"Maths", Class: "jss1"}
	student1.TakeCourse(course1)
	student2.TakeCourse(course1)

	fmt.Println(student1)
	student1.TakeCourse(course3)

	teacher := model.Teacher{Staff:staff, Course: course1}
	teacher2 := model.Teacher{Staff:staff, Course: course3}
	teacher.GradeStudent(student1, 50)
	teacher2.GradeStudent(student1, 60)
	fmt.Println(student1)
	fmt.Println(student2)

	fmt.Println("**************")
	fmt.Println(school.Students)
	fmt.Println("**************")

	school.Expel(*student1, &school.Students)
	fmt.Println(school.Students)
}

