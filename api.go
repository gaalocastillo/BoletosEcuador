package main

import (
  // "fmt"
  "net/http"

  "github.com/gin-gonic/gin"
)

func fetchAllEvents(c *gin.Context){
  var events = []Event{
    Event{1, "Concierto", "no", "Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium", "11-02-2019", "Luis Miguel en Concierto"},
    Event{2, "Obra de Teatro", "no", "Et harum quidem rerum facilis est et expedita distinctio", "21-03-2019", "Sharon Live Action"},
    Event{3, "Concierto", "si", "Ut enim ad minima veniam, quis nostrum exercitationem ullam corporis suscipit laboriosam", "04-01-2020", "La Valicion: Live"},
    Event{4, "Concierto", "no", "Ut enim ad minima veniam, quis nostrum exercitationem ullam corporis suscipit laboriosam", "18-10-2019", "GV Sharks"},
    Event{5, "Ventriloco", "no", " Itaque earum rerum hic tenetur a sapiente delectus, ut aut reiciendis voluptatibus maiores alias consequatu", "14-12-2019", "Quincegrama"},
  }
  c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": events})
}

func fetchAllEventss() []Event{
  var events = []Event{
    Event{1, "Concierto", "no", "Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium", "11-02-2019", "Luis Miguel en Concierto"},
    Event{2, "Obra de Teatro", "no", "Et harum quidem rerum facilis est et expedita distinctio", "21-03-2019", "Sharon Live Action"},
    Event{3, "Concierto", "si", "Ut enim ad minima veniam, quis nostrum exercitationem ullam corporis suscipit laboriosam", "04-01-2020", "La Valicion: Live"},
    Event{4, "Concierto", "no", "Ut enim ad minima veniam, quis nostrum exercitationem ullam corporis suscipit laboriosam", "18-10-2019", "GV Sharks"},
    Event{5, "Ventriloco", "no", " Itaque earum rerum hic tenetur a sapiente delectus, ut aut reiciendis voluptatibus maiores alias consequatu", "14-12-2019", "Quincegrama"},
  }
  return events
}
