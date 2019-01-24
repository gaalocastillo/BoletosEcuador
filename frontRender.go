package main

import (
	"net/http"
  "fmt"
  "encoding/json"
  "io/ioutil"
  "log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func renderEvents(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("user")
  events:= doEventsRequirement(c).Data
  c.HTML(http.StatusOK, "index.tmpl", gin.H{
    "data": events,
    "user": user,
  })
}


func renderEvent(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("user")
	event:= doEventRequirement(c, c.Param("id")).Data

  c.HTML(http.StatusOK, "event.tmpl", gin.H{
    "user": user,
		"event": event,
  })
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


func doEventsRequirement(c *gin.Context) ResponseEvents{
  response, err := http.Get("http://127.0.0.1:3001/api/events")
  fmt.Println("Err: ", err)
  if err != nil || response.StatusCode != http.StatusOK {
    c.Status(http.StatusServiceUnavailable)
  }

  body, readErr := ioutil.ReadAll(response.Body)
  fmt.Println("Response ", string(body))
  if readErr != nil {
    log.Fatal(readErr)
  }
  events := ResponseEvents{}

  jsonErr := json.Unmarshal(body, &events)
  if jsonErr != nil {
    log.Fatal(jsonErr)
  }
  return events
}

func doEventRequirement(c *gin.Context, ID string) ResponseEvent{
  response, err := http.Get("http://127.0.0.1:3001/api/events/" + ID)
  fmt.Println("Err: ", err)
  if err != nil || response.StatusCode != http.StatusOK {
    c.Status(http.StatusServiceUnavailable)
  }

  body, readErr := ioutil.ReadAll(response.Body)
  fmt.Println("Response ", string(body))
  if readErr != nil {
    log.Fatal(readErr)
  }
  event := ResponseEvent{}
  jsonErr := json.Unmarshal(body, &event)
  if jsonErr != nil {
    log.Fatal(jsonErr)
		c.AbortWithStatus(http.StatusNotFound)
  }
  return event
}
