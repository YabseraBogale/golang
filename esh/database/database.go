package database

import (
	"time"
)

type Job struct {
	JobId          int
	JobTitle       string
	JobDescription string
}

func InsertJob(jobtitle string, jobdescription string) Job {
	return Job{
		JobTitle:       jobtitle,
		JobDescription: jobdescription,
	}
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

func InsertEmergencyContact(firstname, middlename, lastname, phonenumber, email, fydaid string) EmergencyContact {
	return EmergencyContact{
		FirstName:   firstname,
		MiddleName:  middlename,
		LastName:    lastname,
		Phonenumber: phonenumber,
		Email:       email,
		FydaId:      fydaid,
	}
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

func InsertCandate(
	firstname string, middlename string, lastname string,
	phonenumber string, email string,
	fydaid string, cv string, hiredate time.Time, jobid int) Candate {
	return Candate{
		FirstName:   firstname,
		MiddleName:  middlename,
		LastName:    lastname,
		Phonenumber: phonenumber,
		Email:       email,
		FydaId:      fydaid,
		HireDate:    hiredate,
		Cv:          cv,
		JobId:       jobid,
	}
}

type Employee struct {
	EmployeeId         int
	FirstName          string
	MiddleName         string
	LastName           string
	Phonenumber        string
	FydaId             string
	Email              string
	HireDate           time.Time
	EmergencyContactId int
	JobId              int
}

func InsertEmployee(
	firstname string, middlename string, lastname string,
	phonenumber string, email string,
	fydaid string, hiredate time.Time, jobid int) Employee {
	return Employee{
		FirstName:   firstname,
		MiddleName:  middlename,
		LastName:    lastname,
		Phonenumber: phonenumber,
		Email:       email,
		FydaId:      fydaid,
		HireDate:    hiredate,
		JobId:       jobid,
	}
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
