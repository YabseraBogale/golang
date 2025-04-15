package database

import "time"

type Job struct {
	JobId          int
	JobTitle       string
	JobDescription string
}

type EmergencyContact struct {
	EmergencyContactId int
	FirstName          string
	MiddleName         string
	LastName           string
	Phonenumber        string
	Email              string
	FydaId             string
}

type Candate struct {
	CandateId   int
	FirstName   string
	MiddleName  string
	LastName    string
	Phonenumber string
	FydaId      string
	Email       string
	HireDate    time.Time
	Cv          string
	JobId       int
}

type Employee struct {
	EmployeeId         int
	FirstName          string
	MiddleName         string
	LastName           string
	Phonenumber        int
	FydaId             string
	Email              string
	HireDate           time.Time
	EmergencyContactId int
	JobId              int
}

type Item struct {
	EmployeeId      int
	ItemId          int
	ItemName        string
	ItemDescription string
	ItemQuantity    int
	ItemStatus      string
	ItemDate        time.Time
}
