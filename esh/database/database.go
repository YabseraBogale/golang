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
	JobTitle    string
}

func InsertCandate(
	firstname string, middlename string, lastname string,
	phonenumber string, email string,
	fydaid string, cv string, hiredate time.Time, jobtitle string) Candate {
	return Candate{
		FirstName:   firstname,
		MiddleName:  middlename,
		LastName:    lastname,
		Phonenumber: phonenumber,
		Email:       email,
		FydaId:      fydaid,
		HireDate:    hiredate,
		Cv:          cv,
		JobTitle:    jobtitle,
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
	JobTitle           string
}

func InsertEmployee(
	firstname string, middlename string, lastname string,
	phonenumber string, email string,
	fydaid string, hiredate time.Time, jobtitle string) Employee {
	return Employee{
		FirstName:   firstname,
		MiddleName:  middlename,
		LastName:    lastname,
		Phonenumber: phonenumber,
		Email:       email,
		FydaId:      fydaid,
		HireDate:    hiredate,
		JobTitle:    jobtitle,
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

func InsertItem(
	employeeid int, itemid int, itemname string,
	itemdescription string, itemquanitiy int,
	itemstatus string, itemdate time.Time) Item {
	return Item{
		EmployeeId:      employeeid,
		ItemId:          itemid,
		ItemName:        itemname,
		ItemDescription: itemdescription,
		ItemQuantity:    itemquanitiy,
		ItemStatus:      itemstatus,
		ItemDate:        itemdate,
	}
}
