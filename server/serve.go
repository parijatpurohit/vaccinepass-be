package server

import (
	"log"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func Serve() {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With"}

	router.Use(cors.New(config))

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
