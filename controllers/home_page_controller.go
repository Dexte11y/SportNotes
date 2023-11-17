package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHomePage(c *gin.Context) {
	c.HTML(http.StatusOK, "homePage.html", nil)
}
