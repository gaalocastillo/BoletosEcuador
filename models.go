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
	  Age          string      `json:"age"`
	  User_type    int         `json:"type"`
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


// Joke contains information about a single Joke
type Joke struct {
	ID     int     `json:"id" binding:"required"`
	Likes  int     `json:"likes"`
	Joke   string  `json:"joke" binding:"required"`
  }
  
  // We'll create a list of jokes
  var jokes = []Joke{
	Joke{1, 0, "Did you hear about the restaurant on the moon? Great food, no atmosphere."},
	Joke{2, 0, "What do you call a fake noodle? An Impasta."},
	Joke{3, 0, "How many apples grow on a tree? All of them."},
	Joke{4, 0, "Want to hear a joke about paper? Nevermind it's tearable."},
	Joke{5, 0, "I just watched a program about beavers. It was the best dam program I've ever seen."},
	Joke{6, 0, "Why did the coffee file a police report? It got mugged."},
	Joke{7, 0, "How does a penguin build it's house? Igloos it together."},
  }
