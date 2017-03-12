package main

import "math/rand"
import "gopkg.in/gin-gonic/gin.v1"

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

func barsPayload(c *gin.Context) {
	c.JSON(200, gin.H{
		"buttons": buttons(),
		"bars":    bars(),
	})
}

func main() {
	r := gin.Default()
	r.GET("/bars", barsPayload)
	r.Run() // listen and serve on 8080
}
