package main

import (
	"net/http"

	"github.com/Ed1123/us-visa-wait-times/usvisa"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.LoadHTMLFiles("./templates/table.tmpl")

	// Define a route to display the table
	router.GET("/table-test", func(c *gin.Context) {
		cities := usvisa.GetWaitData()
		c.HTML(http.StatusOK, "table.tmpl", gin.H{"tittle": "US Visa Wait times", "data": cities})
	})

	router.GET("/wait-times", func(c *gin.Context) {
		cities := usvisa.GetWaitData()
		c.JSON(http.StatusOK, cities)
	})

	// Start the server
	router.Run()
}
