package test

import (
	"github.com/olusolaa/model"
	"testing"
)

func TestAdmit(t testing.T)  {
	staff := model.Staff{Name: "Toluwase", Role:   "principal"}
	principal := model.Principal{Staff: staff, SchoolName: "BestWay"}
	school := model.School{Name: "Best Way", Principal: principal}
	var temp = []struct {
		input    struct{app model.Applicant; sch []model.Student}
		expected int
	}{
		{
			struct {
				app model.Applicant
				sch []model.Student
			}{app: model.Applicant{Name: "Ayo", Age: 8, ClassName: "jss1"}, sch: school.Students}, 0},
{
			struct {
				app model.Applicant
				sch []model.Student
			}{app: model.Applicant{Name: "James", Age: 14, ClassName: "jss1"}, sch: school.Students}, 1},
{
			struct {
				app model.Applicant
				sch []model.Student
			}{app: model.Applicant{Name: "Tobi", Age: 16, ClassName: "jss1"}, sch: school.Students}, 2},
{
			struct {
				app model.Applicant
				sch []model.Student
			}{app: model.Applicant{Name: "Temi", Age: 13, ClassName: "jss1"}, sch: school.Students}, 3},
	}

	for _, test:= range temp{
		if actual:= school.Admit(test.input.app, &test.input.sch); actual != test.expected{
			t.Error("Test failed", test.input, test.expected, actual)
		}
	}
}