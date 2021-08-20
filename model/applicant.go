package model

type Applicant struct {
	schoolName string
	Name string
	Age int
	ClassName string
}

func NewApplicant(age int, className string) *Applicant {
	return &Applicant{Age: age, ClassName: className}
}


func (Applicant) apply(applicant Applicant)  {

}
