package main

import (
  // "fmt"
  "net/http"
  "log"

  "github.com/jinzhu/gorm"
  "github.com/gin-contrib/sessions"
  "github.com/gin-contrib/sessions/redis"
  "github.com/gin-gonic/gin"
   _"github.com/jinzhu/gorm/dialects/postgres"
)


func main() {

  // Set the router as the default one shipped with Gin
  router := gin.Default()
  store, _ := redis.NewStore(64, "tcp", "18.224.37.116:6379", "", []byte("secret"))
  router.Use(sessions.Sessions("mysession", store))

  // Serve frontend static files
  router.LoadHTMLGlob("views/*")
  router.Static("/stylesheets", "./public/stylesheets")
  router.Static("/js", "./public/js")


  router.GET("/events", renderEvents)
  router.GET("/events/:id", renderEvent)
  router.GET("/login", renderLogin)
  router.GET("/checkoutfinish", renderCheckoutFinish)
  router.GET("/profile", renderProfile)
  router.POST("/api/login", login)
  router.GET("/logout", logout)

  router.GET("/", func(c *gin.Context){
    c.Redirect(http.StatusMovedPermanently, "/events")
  })

  // Setup route group for the API
  api := router.Group("/api")
  api.Use(Database())
  {
    api.GET("/events", fetchAllEvents)
    api.GET("/events/:id", fetchSingleEvent)
    api.GET("/seats/:eventID", fetchAvailableSeats)
    api.GET("/user_tickets/", fetchUserTickets)
    api.POST("/purchase", purchaseTickets)
  }

  router.Run(":3001")   // Start and run the server
}

func Database() gin.HandlerFunc {
  const addr = "postgresql://cucaracha@18.224.37.116:26257/boletos_ecuador_db?sslmode=disable"
  db, err := gorm.Open("postgres", addr)
  if err != nil {
      log.Fatal(err)
  }
  // defer db.Close()

  db.AutoMigrate(&VenueModel{})

  return func(c *gin.Context) {
      c.Set("DB", db)
      c.Next()
  }
}
