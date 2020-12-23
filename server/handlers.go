package server

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/parijatpurohit/vaccinepass/server/transport"

	"github.com/parijatpurohit/vaccinepass/logic"
)

func hello(c *gin.Context) {
	c.String(http.StatusOK, "hello")
}

func getUserDetails(c *gin.Context) {
	reqData := &transport.GetUserDetailsRequest{}

	err := c.Bind(reqData)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	res, err := logic.GetClient().GetUserDetails(reqData)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.String(http.StatusOK, toJSON(res))
	return
}

func getUserVaccineDetails(c *gin.Context) {
	reqData := &transport.GetUserVaccineDetailsRequest{}

	err := c.Bind(reqData)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	res, err := logic.GetClient().GetUserVaccineDetails(reqData)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.String(http.StatusOK, toJSON(res))
	return
}

func getRequiredVaccineDetails(c *gin.Context) {
	reqData := &transport.GetRequiredVaccineRequest{}

	err := c.Bind(reqData)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	res, err := logic.GetClient().GetRequiredVaccines(reqData)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.String(http.StatusOK, toJSON(res))
	return
}

func toJSON(data interface{}) string {
	res, _ := json.Marshal(data)
	return string(res)
}
