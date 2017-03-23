package main

import "math/rand"
import "github.com/gin-gonic/gin"

func randInt(min, max int) int {
	return (min + rand.Intn((max+1)-min))
}

func bars() []int {
	size := randInt(2, 5)
	data := []int{}
	for i := 0; i < size; i++ {
		data = append(data, randInt(10, 90))
	}
	return data
}

func buttons() []int {
	return []int{
		randInt(5, 50), randInt(5, 50),
		randInt(-50, -5), randInt(-50, -5),
	}
}

// Route handler
func barsHandler(c *gin.Context) {

	c.JSON(200, gin.H{
		"buttons": buttons(),
		"bars":    bars(),
	})
}

// Custom middleware
func CorsHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

// See https://github.com/gin-gonic/gin
func main() {
	r := gin.Default()
	r.Use(CorsHeaders())
	r.GET("/bars", barsHandler)
	r.Run() // listen and serve on PORT
}
