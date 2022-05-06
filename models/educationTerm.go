package models

import (
	"time"
)

type EducationTerm struct {
	DisplayName string
	ExternalID string
	StartDate	time.Date
	EndDate		time.Date
}