package main

import (
	"log"
	"net/http"
	"strings"
  "fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Thanks to otraore for the code example
// https://gist.github.com/otraore/4b3120aa70e1c1aa33ba78e886bb54f3

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("user")
		if user == nil {
			// You'd normally redirect to login page
      fmt.Println("Err: ", user)
			c.Redirect(http.StatusMovedPermanently, "/login?statusCode=3")
		} else {
			// Continue down the chain to handler etc
			c.Next()
		}
	}
}

func login(c *gin.Context) {
	session := sessions.Default(c)
	username := c.PostForm("username")
	password := c.PostForm("password")

	if strings.Trim(username, " ") == "" || strings.Trim(password, " ") == "" {
		c.Redirect(http.StatusMovedPermanently, "/login?statusCode=1")
		return
	}
	if (username == "Leonardo" && password == "leonardo") || (username == "Galo" && password == "galo") {
		session.Set("user", username) //In real world usage you'd set this to the users ID
		err := session.Save()
		if err != nil {
			c.Redirect(http.StatusMovedPermanently, "/login?statusCode=2")
		} else {
			c.Redirect(http.StatusMovedPermanently, "/events")
		}
	} else {
		c.Redirect(http.StatusMovedPermanently, "/login?statusCode=1")
	}
}

func logout(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("user")
  fmt.Println("Err: ", user)
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session token"})
	} else {
		fmt.Println("ENTRE ")
		log.Println(user)
		session.Delete("user")
		session.Save()
		c.JSON(200, gin.H{})
	}
}

func privateG(c *gin.Context) {
  session := sessions.Default(c)
  user := session.Get("user")
	c.JSON(http.StatusOK, gin.H{"hello": user})
}
