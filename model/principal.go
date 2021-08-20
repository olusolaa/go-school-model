package model

import (
	"fmt"
)

type Principal struct {
	Staff
	SchoolName string
}

func (principal *Principal) Admit(applicant Applicant, students *[]Student) int {
	if principal.Name == ""{
		fmt.Println("This school does not have a principal, Admission can not take place")
		return len(*students)
	}
	if applicant.Age >= 12 {
		newStudent := Student{applicant, len(*students)+1, make([]Course, 0)}
		*students = append(*students, newStudent)
	}else {
		fmt.Printf("%s, not admitted", applicant.Name)
	}
	return len(*students)
}

func (Principal) addTeacher(staff Staff, course Course, teacherList *[]Teacher)  {
	if staff.Role == "Teacher"{
		*teacherList = append(*teacherList, Teacher{staff, course})
	}else {
		fmt.Printf("an error occured. Expecting a teaching staff, recieved a %s staff\n", staff.Role)
	}
}

func (Principal) Expel(student Student, studentList *[]Student)  {
	studentGPA:= CalculateGPA(student)
	if studentGPA <50{
		for i, student1 := range *studentList {
			if student1.ID == student.ID{
				list := *studentList
				list = append(list[:i], list[i+1:]...)
				*studentList = list
				return
			}
		}
		fmt.Printf("%s record not found, cannot be expeled", student.Name)
	}
}



