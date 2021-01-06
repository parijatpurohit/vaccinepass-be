package server

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/parijatpurohit/vaccinepass/server/transport"

	"github.com/parijatpurohit/vaccinepass/logic"
)

func hello(c *gin.Context) {
	c.String(http.StatusOK, "hello")
}

func getUserSummary(c *gin.Context) {
	reqData := &transport.GetUserSummaryRequest{}

	err := c.Bind(reqData)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	res, err := logic.GetClient().GetUserSummary(reqData)
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

func getCountryVaccineDetails(c *gin.Context) {
	reqData := &transport.GetCountryVaccinesRequest{}

	err := c.Bind(reqData)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	res, err := logic.GetClient().GetCountryVaccines(reqData)
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
