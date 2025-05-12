package database

import (
	"time"
)

type Job struct {
	JobId          int
	JobTitle       string
	Office         string
	Department     string
	JobDescription string
}

func InsertJob(jobtitle string, office string, department string, jobdescription string) Job {
	return Job{
		JobTitle:       jobtitle,
		Office:         office,
		Department:     department,
		JobDescription: jobdescription,
	}
}

type Employee struct {
	EmployeeId           int
	FirstName            string
	MiddleName           string
	LastName             string
	Phonenumber          string
	FydaId               string
	Email                string
	HireDate             time.Time
	Salary               int
	YearExperince        int
	EmergencyFirstName   string
	EmergencyMiddleName  string
	EmergencyLastName    string
	EmergencyPhonenumber string
	EmergencyEmail       string
	EmergencyFydaId      string
	Department           string
	Office               string
	JobTitle             string
}

func InsertEmployee(
	firstname string, middlename string, lastname string,
	phonenumber string, email string,
	fydaid string, hiredate time.Time, salary int, year_experince int,
	emergencyfirstname string, emergencymiddlename string, emergencylastname string,
	emergencyphonenumber string, emergencyemail string,
	emergencyfydaid string, department string, office string, jobtitle string) Employee {
	return Employee{
		FirstName:            firstname,
		MiddleName:           middlename,
		LastName:             lastname,
		Phonenumber:          phonenumber,
		Email:                email,
		FydaId:               fydaid,
		HireDate:             hiredate,
		Salary:               salary,
		YearExperince:        year_experince,
		EmergencyFirstName:   emergencyfirstname,
		EmergencyMiddleName:  emergencymiddlename,
		EmergencyLastName:    emergencylastname,
		EmergencyPhonenumber: emergencyphonenumber,
		EmergencyEmail:       emergencyemail,
		EmergencyFydaId:      emergencyfydaid,
		Department:           department,
		Office:               office,
		JobTitle:             jobtitle,
	}
}

type Item struct {
	EmployeeId      string
	ItemId          string
	ItemName        string
	ItemDescription string
	ItemQuantity    float32
	ItemStatus      string
	ItemDate        time.Time
}

func InsertItem(
	employeeid string, itemid string, itemname string,
	itemdescription string, itemquanitiy float32,
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

type PurchaseRequest struct {
	ItemId              string
	EmployeeId          string
	ItemName            string
	ItemDescription     string
	ItemQuantity        int
	ItemStatus          string
	ItemPurchaseRequest string
	ItemDate            time.Time
}

func InsertPurchaseRequest(
	employeeid string, itemid string, itemname string,
	itemdescription string, itemquanitiy int,
	itemstatus string, itempurchaserequest string, itemdate time.Time) PurchaseRequest {
	return PurchaseRequest{
		ItemId:          itemid,
		EmployeeId:      employeeid,
		ItemName:        itemname,
		ItemDescription: itemdescription,
		ItemQuantity:    itemquanitiy,
		ItemStatus:      itemstatus,
		ItemDate:        itemdate,
	}
}
