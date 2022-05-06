package models

import (
	
)

type EducationClass struct {
	Id		string
	DisplayName		string
	MailNickname	string
	Description		string
	CreatedBy	EducationUser
	ClassCode	string
	ExternalName	string
	ExternalID	string
	Grade	string
	Term 	EducationTerm
}