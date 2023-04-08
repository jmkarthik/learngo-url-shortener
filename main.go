package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println(GetURL("goog"))
	r := gin.Default()

	r.GET("/l/:shortCode", func(c *gin.Context) {
		shortCode := c.Params.ByName("shortCode")
		url, foundURL := GetURL(shortCode)
		if foundURL == true {
			c.Redirect(http.StatusMovedPermanently, url)
		} else {

			c.AbortWithError(http.StatusNotFound, nil)
		}
	})
	r.Run()

}
