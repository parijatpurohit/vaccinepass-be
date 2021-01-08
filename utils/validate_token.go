package utils

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
)

func ValidateToken(c *gin.Context) error {
	auth := strings.TrimSpace(c.GetHeader("Authorization"))
	if auth == "" || auth == "Bearer" {
		return errors.New("random error which i just made up")
	}
	return nil
}
