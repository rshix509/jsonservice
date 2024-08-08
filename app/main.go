package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	Utils "github.com/rshix509/jsonservice/app/lib"
)

var F Utils.FileInMem

func init() {
	F = Utils.FileInMem{
		Filename: "../large-file.json",
	}
	F.ReadContentsAndStoreStruct()
}

func main() {
	fmt.Println("BLAH")
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.GET("/health/check", Pingpong)
		v1.GET("/stream", GetStream)
	}

	router.Run()
}
