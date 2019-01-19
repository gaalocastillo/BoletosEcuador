package main

import (
  // "fmt"
  "net/http"
  "strconv"
  //"database/sql"
  "log"

  "github.com/jinzhu/gorm"
  "github.com/gin-gonic/contrib/sessions"
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

  store := sessions.NewCookieStore([]byte("secret"))
  router.Use(sessions.Sessions("mysession", store))

  router.LoadHTMLGlob("views/*")
  // Serve frontend static files
  router.Static("/stylesheets", "./public/stylesheets")
  router.Static("/js", "./public/js")

  private := router.Group("/private", AuthRequired())
  private.GET("/", privateG)

  router.POST("/api/login", login)
  router.GET("/api/logout", logout)

  router.GET("/events", renderEvents)
  router.GET("/events/:id", renderEvent)
  router.GET("/login", renderLogin)
  router.GET("/checkoutfinish", renderCheckoutFinish)
  router.GET("/profile", renderProfile)

  router.GET("/", func(c *gin.Context){
    c.Redirect(http.StatusMovedPermanently, "/events")
  })

  // router.GET("/events", func(c *gin.Context){
  //   c.HTML(http.StatusOK, "index.tmpl", gin.H{
  //     "title": "Titulo del Sitio",
  //     "data": events,
  //     "user": "not defined",
  //   })
  // })

  // router.GET("/events/:id", func(c *gin.Context){
  //   if eventId, err := strconv.Atoi(c.Param("id")); err == nil {
  //     for i := 0; i < len(events); i++ {
  //       if events[i].ID == eventId {
  //         c.HTML(http.StatusOK, "event.tmpl", gin.H{
  //           "event": events[i],
  //           "user": "not defined",
  //         })
  //       }
  //     }
  //   } else {
  //     c.AbortWithStatus(http.StatusNotFound)
  //   }
  // })

  // router.GET("/profile", func(c *gin.Context){
  //   c.HTML(http.StatusOK, "index.tmpl", gin.H{
  //     "title": "Titulo del Sitio",
  //     "user": "Pepito Pihuave",
  //   })
  // })

  // router.GET("/login", func(c *gin.Context){
  //   c.HTML(http.StatusOK, "login.tmpl", gin.H{
  //     "title": "Titulo del Sitio",
  //     "user": "Pepito Pihuave",
  //   })
  // })

  // router.GET("/checkoutfinish", func(c *gin.Context){
  //   errCode := c.Query("err")
  //   // fmt.Println("Err: ", errCode)
  //   c.HTML(http.StatusOK, "checkoutfinish.tmpl", gin.H{
  //     "title": "Titulo del Sitio",
  //     "errCode": errCode,
  //     "user": "Pepito Pihuave",
  //   })
  // })

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
    api.GET("/seats/:eventID", fetchAvailableSeats)
    api.GET("/user_tickets/:userID", fetchUserTickets)
    api.POST("/purchase", purchaseTickets)
  }


  // Start and run the server
  router.Run(":3001")
}
