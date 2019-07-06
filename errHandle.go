package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func errHandle(err error, c *gin.Context) {
	log.Println(err)
	c.JSON(400, gin.H{
		"errMessage": err,
	})
}
