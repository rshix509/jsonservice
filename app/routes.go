package main

import (
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var BYTE_SIZE = 20000

func Pingpong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"ping": "pong",
	})
}

func GetStream(c *gin.Context) {
	iop := F.ReadContentsByPart()
	c.Stream(func(w io.Writer) bool {
		data := make([]byte, BYTE_SIZE)
		_, err := iop.Read(data)
		if err != nil && err.Error() != "EOF" {
			log.Println("Error occurred while reading contents" + err.Error())
			return false
		}

		if err != nil && err.Error() == "EOF" {
			log.Println("Done reading data")
			return false
		}
		c.SSEvent("message", string(data))
		return true

	})
}
