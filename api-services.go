package main
import (
	"net/http"
	"strconv"
	//"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/gin-gonic/gin"
	 _"github.com/jinzhu/gorm/dialects/postgres"

)

/**
API web services functions
*/
// fetchAllEvents fetches all events
func fetchAllEvents(c *gin.Context) {
	var events []EventModel
	db := c.MustGet("DB").(*gorm.DB)
  db.Find(&events)
	if len(events) <= 0 {
	  c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No event found!"})
	  return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": events})
}

  // fetchSingleEvent fetches a single event
func fetchSingleEvent(c *gin.Context) {
	var event EventModel
  eventID := c.Param("id")
	db := c.MustGet("DB").(*gorm.DB)
	db.First(&event, eventID)
	if event.ID == 0 {
	  c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No event found!"})
	  return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": event})
}

// fetchAvailableSeats obtains all seats available, given an event
func fetchAvailableSeats(c *gin.Context) {
	var event EventModel
	var venue VenueModel
	var zones []ZoneModel
	var allSeats []SeatModel
	var availableSeats []DummySeat

	eventID := c.Param("eventID")
	db := c.MustGet("DB").(*gorm.DB)
	db.First(&event, eventID)
	if event.ID == 0 {
		fmt.Println("ALGO PASA")
	  c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No event found!"})
	  return
	}
	db.Model(&event).Related(&venue)
	db.Model(&venue).Related(&zones)
	for _, zone := range zones {
			db.Model(&zone).Related(&allSeats)
			for _, seat := range allSeats {
				if seat.IsAvailable == true{
					ds := DummySeat{int(seat.ID), seat.Number, zone.Name, zone.Price}
					availableSeats = append(availableSeats, ds)
				}
			}
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": availableSeats})
	return
}

func fetchUserTickets(c *gin.Context) {
	var userTickets []TicketModel
	var user UserModel

	userID := 1
	db := c.MustGet("DB").(*gorm.DB)
	db.First(&user, userID)
	if userID <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "User not found!"})
		return
	}
	db.Model(&user).Related(&userTickets)
	var tempSeat SeatModel
	var tempZone ZoneModel
	var tempEvent EventModel
	var data []DummyTicket
	for _, ticket := range userTickets{
		db.Model(&ticket).Related(&tempSeat)
		db.Model(&tempSeat).Related(&tempZone)
		db.Model(&ticket).Related(&tempEvent)
		dt := DummyTicket{int(ticket.ID), ticket.Number, tempSeat.Number, tempZone.Name, tempZone.Price, tempEvent.Name}
		data = append(data, dt)
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": data})
	return
}

   // insert new tickets purchase
func purchaseTickets(c *gin.Context) {
	seatsAmount, _ := strconv.Atoi(c.PostForm("seats"))
	userID, _ := strconv.Atoi(c.PostForm("user-ID"))
	eventID, _ := strconv.Atoi(c.PostForm("event-ID"))
	seatsIds := make([]int, seatsAmount)
	eventID = eventID +1
	userID = userID +1

	for i := 0; i < seatsAmount; i++ {
	  //ticketsIds[i] = strconv. c.PostForm("tickets-ids")[i]
	  // if jokes[i].ID == jokeid {
	  //  jokes[i].Likes += 1
	 // }
	}
  //  ticketsIds = c.PostForm("tickets-ids")
	fmt.Println(seatsIds)
  //  completed, _ := strconv.Atoi(c.PostForm("completed"))
   // todo := todoModel{Title: c.PostForm("title"), Completed: completed}
  //  db.Save(&todo)
	//c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Todo item created successfully!", "resourceId": todo.ID})
}
