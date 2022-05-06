package models

import (
	"time"
)

type EducationStudent struct {
	BirthDate 	time.Date
	ExternalID	string
	Gender	string
	Grade	string
	GraduationYear	string
	StudentNumber	string
}