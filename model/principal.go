package model

import "fmt"

type Principal struct {
	Staff
	SchoolName string
}

func (principal *Principal) Admit(applicant Applicant, students *[]Student)  {
	if principal.Name == ""{
		fmt.Println("This school does not have a principal, Admission can not take place")
		return
	}
	if applicant.Age >= 12 {
		newStudent := Student{applicant, len(*students)+1, make([]Course, 0)}
		*students = append(*students, newStudent)
	}else {
		fmt.Printf("%s, not admitted", applicant.Name)
	}
}

func (Principal) addTeacher(staff Staff, course Course, teacherList *[]Teacher)  {
	if staff.Role == "Teacher"{
		*teacherList = append(*teacherList, Teacher{staff, course})
		course.TeacherName = staff.Name
	}else {
		fmt.Printf("an error occured. Expecting a teaching staff, recieved a %s\n", staff.Role)
	}
}


