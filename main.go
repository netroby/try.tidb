package main

import "gopkg.in/gin-gonic/gin.v1"

func main() {
	r := gin.Default()
	r.Static("/javascripts", "./public/javascripts")
	r.Static("/images", "./public/images")
	r.StaticFile("/", "./public/index.html")
	r.StaticFile("/style.css", "./public/style.css")
	r.GET("/eval", func(c *gin.Context) {
		//when success
		//c.JSON(200, gin.H{"response": "80", "session_id": "24e2c503a3976d63282abaa192629aeee0a5944eb8679a6e23247a65c4d2b962"})
		//when failed
		c.JSON(200, gin.H{"error": "error", "session_id": "24e2c503a3976d63282abaa192629aeee0a5944eb8679a6e23247a65c4d2b962"})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
