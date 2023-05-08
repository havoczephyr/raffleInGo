package main

import (
	"strconv"
)

var lastTicketNumber = 100

type Entry struct {
	firstName    string
	lastName     string
	telephone    string
	ticketNumber int
}

func (e *Entry) SetValues(firstName, lastName, telephone string) {
	e.firstName = firstName
	e.lastName = lastName
	e.telephone = telephone

	lastTicketNumber++
	e.ticketNumber = lastTicketNumber
}

func (e *Entry) GetName() string {
	return e.lastName + ", " + e.firstName
}

func (e *Entry) GetTelephone() string { return e.telephone }

func (e *Entry) GetTicketNumber() int { return e.ticketNumber }

func (e *Entry) TicketData() string {
	return e.GetName() + " Phone: " + e.GetTelephone() + " Ticket No.: " + strconv.Itoa(e.GetTicketNumber())
}
