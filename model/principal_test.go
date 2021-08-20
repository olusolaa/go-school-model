package model
//
//import "testing"
//
//func TestPrincipal_Admit(t *testing.T) {
//	type fields struct {
//		Staff      Staff
//		SchoolName string
//	}
//	type args struct {
//		applicant Applicant
//		students  *[]Student
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//		want   int
//	}{
//		{"MODEL"
//		struct {
//				app model.Applicant
//				sch []model.Student
//			}{app: model.Applicant{Name: "Ayo", Age: 8, ClassName: "jss1"}, sch: school.Students}, 0},
//		{
//			struct {
//				app model.Applicant
//				sch []model.Student
//			}{app: model.Applicant{Name: "James", Age: 14, ClassName: "jss1"}, sch: school.Students}, 1},
//		{
//			struct {
//				app model.Applicant
//				sch []model.Student
//			}{app: model.Applicant{Name: "Tobi", Age: 16, ClassName: "jss1"}, sch: school.Students}, 2},
//		{
//			struct {
//				app model.Applicant
//				sch []model.Student
//			}{app: model.Applicant{Name: "Temi", Age: 13, ClassName: "jss1"}, sch: school.Students}, 3},
//	}
//
//	// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			principal := &Principal{
//				Staff:      tt.fields.Staff,
//				SchoolName: tt.fields.SchoolName,
//			}
//			if got := principal.Admit(tt.args.applicant, tt.args.students); got != tt.want {
//				t.Errorf("Admit() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
