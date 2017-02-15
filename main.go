package main

import (
	"fmt"
	"os/exec"

	"gopkg.in/gin-gonic/gin.v1"
)

func main() {
	r := gin.Default()
	r.Static("/javascripts", "./public/javascripts")
	r.Static("/images", "./public/images")
	r.StaticFile("/", "./public/index.html")
	r.StaticFile("/style.css", "./public/style.css")
	r.GET("/eval", func(c *gin.Context) {
		sessionID := c.Query("session_id")
		cmd := c.Query("command")
		out, err := exec.Command("/usr/bin/mysql", "-h", "tidb", "-P", "4000", "-D", "test", "-e", cmd).Output()
		if err != nil {
			c.JSON(200, gin.H{"error": err.Error(), "session_id": sessionID})
		}
		outstr := string(out)
		fmt.Println(outstr)
		c.JSON(200, gin.H{"response": outstr, "session_id": sessionID})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
