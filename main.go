package main

import (
  "net/http"
  "strconv"
  //"database/sql"
  "log"

  "github.com/jinzhu/gorm"  
  "github.com/gin-gonic/gin"
   _"github.com/jinzhu/gorm/dialects/postgres"
  
)

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
    api.GET("/seats/eventID")
    api.POST("/purchase", purchaseTickets)
  }

  api.GET("/jokes", JokeHandler)
  api.POST("/jokes/like/:jokeID", LikeJoke)

  // Start and run the server
  router.Run(":3001")
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
