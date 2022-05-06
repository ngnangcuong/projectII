package models

import (
	"time"
)

type EducationUser struct {
	AccountEnabled	bool
	BusinessPhone	[]string
	CreatedBy	string
	Department	string
	DisplayName	string
	GivenName	string
	Id	string
	Mail	string
	MailNickname	string
	MiddleName	string
	SurName	string
	PasswordProfile	string
	PrimaryRole	string
	Student		EducationStudent
	Teacher		EducationTeacher
	UserType	string
}