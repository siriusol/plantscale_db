package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	register(r)

	Init()

	r.Run(":8081")
}
