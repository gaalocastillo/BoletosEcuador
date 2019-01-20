package main

import (
  "net/http"
  "strconv"
  //"database/sql"
  "log"
  "fmt"

  "github.com/jinzhu/gorm"  
  "github.com/gin-gonic/gin"
   _"github.com/jinzhu/gorm/dialects/postgres"
  
)


// Joke contains information about a single Joke
type Joke struct {
  ID     int     `json:"id" binding:"required"`
  Likes  int     `json:"likes"`
  Joke   string  `json:"joke" binding:"required"`
}

type (
  VenueModel struct {
    //un venue tiene varias zonas
    gorm.Model
    VenueType   string      `json:"type"`
    Name         string      `json:"name"`
    Country      int         `json:"country"`
    City         int         `json:"city"`
    Address      int         `json:"address"`
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
    Price        float64     `json:"country"`
    VenueModelID uint         `json:"venuemodelid"`  
    Seats []SeatModel  
   }

   SeatModel struct {
    //una zona tiene varios asientos
    gorm.Model
    Number       int         `json:"number"`
    IsAvailable  int         `json:"isavailable"`
    ZoneModelID  uint        `json:"zonemodelid"`  
   }

   UserModel struct {
    //un usuario tiene varios tickets
    gorm.Model
    Username     string      `json:"username"`
    Age          string      `json:"age"`
    User_type         int    `json:"type"`
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

//function that retrieves a certain venue

func main() {
  //establish connection with DB
  /*db, err := sql.Open("postgres",
        "postgresql://cucaracha:cucarachaAdmin@localhost:26257/boletos_ecuador_db?ssl=true&sslmode=require&sslrootcert=certs/ca.crt&sslkey=certs/client.cucaracha.key&sslcert=certs/client.cucaracha.crt")
  if err != nil {
        log.Fatal("error connecting to the database: ", err)
  }
  defer db.Close()*/
  const addr = "postgresql://cucaracha:cucarachaAdmin@localhost:26257/boletos_ecuador_db?ssl=true&sslmode=require&sslrootcert=certs/ca.crt&sslkey=certs/client.cucaracha.key&sslcert=certs/client.cucaracha.crt"
  db, err := gorm.Open("postgres", addr)
  if err != nil {
      log.Fatal(err)
  }
  defer db.Close()

  db.AutoMigrate(&VenueModel{})

  // Set the router as the default one shipped with Gin
  router := gin.Default()
  router.LoadHTMLGlob("views/*")
  router.Static("/stylesheets", "./public/stylesheets")
  // Serve frontend static files

  router.GET("/", func(c *gin.Context){
    c.HTML(http.StatusOK, "index.tmpl", gin.H{
      "title": "Titulo del Sitio",
      "user": "not defined",
    })
  })

  router.GET("/events:id", func(c *gin.Context){
    c.HTML(http.StatusOK, "index.tmpl", gin.H{
      "title": "Titulo del Sitio",
    })
  })

  /*router.GET("/venues/:id", func(c *gin.Context){
    var venue_name string
    var address string
    var city string
    var country string
    var venue_type string

    if venue_id, err := strconv.Atoi(c.Param("id")); err == nil {
      // find venue and print name
      row := db.QueryRow("SELECT name,address,city,country,type FROM boletos_ecuador_db.venue WHERE id=$1",venue_id)
      switch err := row.Scan(&venue_name, &address, &city, &country, &venue_type); err {
        case sql.ErrNoRows:
          fmt.Println("No rows were returned!")
        case nil:
          c.HTML(http.StatusOK, "venue.tmpl", gin.H{
            "venue_name": venue_name,
            "address": address,
            "city":city,
            "country":country,
            "type":venue_type,
      "user": "not defined",
            
          })
        default:
          panic(err)

      }

    } else {
      // Joke ID is invalid
      c.AbortWithStatus(http.StatusNotFound)
    }
  })
  */

  router.GET("/profile", func(c *gin.Context){
    c.HTML(http.StatusOK, "index.tmpl", gin.H{
      "title": "Title",
    })
  })

  router.GET("/checkoutfinish", func(c *gin.Context){
    c.HTML(http.StatusOK, "index.tmpl", gin.H{
      "title": "Titulo del Sitio",
    })
  })

  router.GET("/checkout", func(c *gin.Context){
    c.HTML(http.StatusOK, "index.tmpl", gin.H{
      "title": "Titulo del Sitio",
    })
  })

  router.GET("/events", func(c *gin.Context){
    c.HTML(http.StatusOK, "index.tmpl", gin.H{
      "title": "Titulo del Sitio",
      "data": jokes,
      "user": "not defined",
    })
  })

  // router.Use(static.Serve("/", static.LocalFile("./views", true)))

  // Setup route group for the API
  api := router.Group("/api")
  {
    api.GET("/", func(c *gin.Context) {
      c.JSON(http.StatusOK, gin.H {
        "message": "pong",
      })
    })

    api.GET("/events", fetchAllEvents)
    api.GET("/events/:id", fetchSingleEvent)
    api.POST("/purchase", purchaseTickets)
  }

  api.GET("/jokes", JokeHandler)
  api.POST("/jokes/like/:jokeID", LikeJoke)

  // Start and run the server
  router.Run(":3001")
}

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

// JokeHandler retrieves a list of available jokes
func JokeHandler(c *gin.Context) {
  c.Header("Content-Type", "application/json")
  c.JSON(http.StatusOK, jokes)
}

// LikeJoke increments the likes of a particular joke Item
func LikeJoke(c *gin.Context) {
  if jokeid, err := strconv.Atoi(c.Param("jokeID")); err == nil {
    // find joke, and increment likes
    for i := 0; i < len(jokes); i++ {
      if jokes[i].ID == jokeid {
        jokes[i].Likes += 1
      }
    }

    // return a pointer to the updated jokes list
    c.JSON(http.StatusOK, &jokes)
  } else {
    // Joke ID is invalid
    c.AbortWithStatus(http.StatusNotFound)
  }
}
