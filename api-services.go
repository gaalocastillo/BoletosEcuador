package main
import (
	"net/http"
	"strconv"
	//"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	 _"github.com/jinzhu/gorm/dialects/postgres"
	
  )

/**
API web services functions
*/
// fetchAllEvents fetches all events
func fetchAllEvents(c *gin.Context) {
	var events []EventModel
  //  var _todos []transformedTodo
  //  db.Find(&events)
	if len(events) <= 0 {
	  c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No event found!"})
	  return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": events})
   }
  
  // fetchSingleEvent fetches a single event
  func fetchSingleEvent(c *gin.Context) {
	var event EventModel
  //  eventID := c.Param("id")
  //  db.First(&todo, todoID)
	if event.ID == 0 {
	  c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No event found!"})
	  return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": event})
   }
  
   // insert new tickets purchase
  func purchaseTickets(c *gin.Context) {  
	seatsAmount, _ := strconv.Atoi(c.PostForm("seats-amount"))
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