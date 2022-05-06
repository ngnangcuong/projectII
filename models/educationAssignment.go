package models

import (
	"time"
)

type EducationAssignment struct {
	Id	string
	AddedStudentAction	string
	AllowLateSubmissions	bool
	AllowStudentsToAddResourcesToSubmission 	bool
	AssignDateTime	time.Time
	AssignTo	EducationAssignmentRecipient
	ClassId	string
	CloseDateTime	time.Time
	CreatedBy	EducationUser
	CreatedDateTime		time.Time
	DisplayName		string
	DueDateTime		time.Time
	Instructions	[]byte
	LastModifiedBy	EducationUser
	LastModifiedDateTime	time.Time
	NotificationChannelUrl	string
	Status 	string
	WebUrl	string
	ResourcesFolderUrl	string
}