package main

import (
	"net/http"
  "strconv"
  "fmt"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func renderEvents(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("user")
  c.HTML(http.StatusOK, "index.tmpl", gin.H{
    "data": events,
    "user": user,
  })
}


func renderEvent(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("user")
  if eventId, err := strconv.Atoi(c.Param("id")); err == nil {
    for i := 0; i < len(events); i++ {
      if events[i].ID == eventId {
        c.HTML(http.StatusOK, "event.tmpl", gin.H{
          "event": events[i],
          "user": user,
        })
      }
    }
  } else {
    c.AbortWithStatus(http.StatusNotFound)
  }
}

func renderLogin(c *gin.Context) {
  session := sessions.Default(c)
  user := session.Get("user")
  statusCode := c.Query("statusCode")
  fmt.Println("Err: ", user)
  if user == nil {
    c.HTML(http.StatusOK, "login.tmpl", gin.H{
      "user": user,
      "statusCode": statusCode,
    })
  } else {
    c.Redirect(http.StatusMovedPermanently, "/events")
  }
}

func renderCheckoutFinish(c *gin.Context) {
  session := sessions.Default(c)
  user := session.Get("user")
  if user == nil {
    c.Redirect(http.StatusMovedPermanently, "/login?statusCode=3")
  } else {
    errCode := c.Query("err")
    c.HTML(http.StatusOK, "checkoutfinish.tmpl", gin.H{
      "errCode": errCode,
      "user": user,
    })
  }
}

func renderProfile(c *gin.Context) {
  session := sessions.Default(c)
  user := session.Get("user")
  if user == nil {
    c.Redirect(http.StatusMovedPermanently, "/login?statusCode=2")
  } else {
    errCode := c.Query("err")
    c.HTML(http.StatusOK, "profile.tmpl", gin.H{
      "errCode": errCode,
      "user": user,
    })
  }
}
