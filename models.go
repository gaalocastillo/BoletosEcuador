package main

import (
	"github.com/jinzhu/gorm"
)

type (
	VenueModel struct {
	  //un venue tiene varias zonas
	  gorm.Model
	  VenueType     string      `json:"type"`
	  Name          string      `json:"name"`
	  Country       string      `json:"country"`
	  City          string      `json:"city"`
	  Address       string      `json:"address"`
	  Zones []ZoneModel
	}

	EventModel struct {
	  //un evento pertenece a un venue
	  gorm.Model
	  Name         string      `json:"name"`
	  Date         string      `json:"date"`
	  Description  string      `json:"description"`
	  EventType    string      `json:"string"`
	  IsSoldOut    bool        `json:"issoldout"`
	  VenueModelID uint        `json:"venuemodelid"`
	  VenueModel   VenueModel  `json:"venuemodel"`
	}

	ZoneModel struct {
	  //un venue tiene varias zonas
	  //una zona tiene varios asientos
	  gorm.Model
	  Name         string      `json:"name"`
	  Price        float64     `json:"price"`
	  VenueModelID uint         `json:"venuemodelid"`
	  Seats []SeatModel
	}

	SeatModel struct {
	  //una zona tiene varios asientos
	  gorm.Model
	  Number       int         `json:"number"`
	  IsAvailable  bool        `json:"isavailable"`
	  ZoneModelID  uint        `json:"zonemodelid"`
	}

	UserModel struct {
	  //un usuario tiene varios tickets
	  gorm.Model
	  Username     string      `json:"username"`
	  Age          int      `json:"age"`
	  User_type    string         `json:"type"`
	  Password     string      `json:"password"`
	  Tickets []TicketModel
	}

	TicketModel struct {
	  //Un ticket pertenece a un usuario
	  //Un ticket tiene un asiento
	  //Un ticket pertenece a un evento
	  gorm.Model
	  Number       int         `json:"number"`
	  SeatModelID  uint        `json:"seatmodelid"`
	  SeatModel    SeatModel   `json:"seatmodel"`
	  UserModelID  uint        `json:"usermodelid"`
	  EventModel   EventModel  `json:"eventmodel"`
	  EventModelID uint        `json:"eventmodelid"`
	}
)

type ResponseEvents struct{
	Data    []EventModel   	`json:"data"`
	Status  int       			`json:"status"`
}

type ResponseEvent struct{
	Data    EventModel   	`json:"data"`
	Status  int       		`json:"status"`
}

  // Joke contains information about a single Joke
type DummySeat struct {
	ID     		string     `json:"id" binding:"required"`
	Number  	int     `json:"number" binding:"required"`
	ZoneName   	string  `json:"zoneName" binding:"required"`
	ZonePrice   float64	`json:"zonePrice" binding:"required"`
}
  // Joke contains information about a single Joke
type DummyTicket struct {
	ID     			string     `json:"id" binding:"required"`
	TicketNumber	int		`json:"ticketNumber" binding:"required"`
	SeatNumber  	int     `json:"seatNumber" binding:"required"`
	ZoneName   		string  `json:"zoneName" binding:"required"`
	ZonePrice   	float64	`json:"zonePrice" binding:"required"`
	EventName   	string	`json:"eventName" binding:"required"`
}
