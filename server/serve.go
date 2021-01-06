package server

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Serve() {
	router := gin.Default()

	router.GET("/hello", hello)
	router.POST("/user-summary", getUserSummary)
	router.POST("/user-vaccines", getUserVaccineDetails)
	router.POST("/country-vaccines", getCountryVaccineDetails)
	err := router.Run(":8090")
	if err != nil {
		panic(err)
	}
	log.Println("Started server on port 8090")
}
