package models

import (
	time.Time
)

type EducationSubmission struct {
	Id string
	Recipient //
	ReturnedBy	EducationUser
	ReturnedDateTime	time.Date
	ResourcesFolderUrl string
	Status	string
	SubmittedBy	EducationUser
	SubmittedDateTime	time.Date
	UnsubmitedBy EducationUser
	UnsubmitedDateTime	time.Date
	ReassignedBy EducationUser
	ReassignedDateTime	time.Date
}