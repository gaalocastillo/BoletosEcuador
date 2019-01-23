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
					ds := DummySeat{strconv.Itoa(int(seat.ID)), seat.Number, zone.Name, zone.Price}
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
		dt := DummyTicket{strconv.Itoa(int(ticket.ID)), ticket.Number, tempSeat.Number, tempZone.Name, tempZone.Price, tempEvent.Name}
		data = append(data, dt)
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": data})
	return
}

// insert new tickets purchase
func purchaseTickets(c *gin.Context) {
	seatID, _ := strconv.Atoi(c.PostForm("seat-ID"))
	db := c.MustGet("DB").(*gorm.DB)
	var seat SeatModel
	db.First(&seat, seatID)
	if seat.ID <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No seat found!"})
		return
	}
	if seat.IsAvailable == false {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusBadRequest, "message": "Sorry. That seat is unavailable!"})
		return
	}

	eventID, _ := strconv.Atoi(c.PostForm("event-ID"))
	userID, _ := strconv.Atoi(c.PostForm("user-ID"))
	var event EventModel
	var ticket TicketModel
	db.First(&event, eventID)
	if userID <= 0{
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No user found!"})
		return
	}
	if event.ID <= 0{
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No event found!"})
		return
	}
	ticket = TicketModel{SeatModel: seat, UserModelID: uint(userID), EventModel: event}
	db.Save(&ticket)
	db.Model(&seat).Update("IsAvailable", false)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Transaction successfully completed!"})
	return
}
