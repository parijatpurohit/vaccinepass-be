package server

import (
	"github.com/gin-gonic/gin"
	"log"
)

func Serve() {
	router := gin.Default()

	router.GET("/hello", hello)
	router.POST("/user", getUserDetails)
	router.POST("/user-vaccine", getUserVaccineDetails)
	router.POST("/country-vaccine", getRequiredVaccineDetails)
	err := router.Run(":8090")
	if err != nil {
		panic(err)
	}
	log.Println("Started server on port 8090")
}