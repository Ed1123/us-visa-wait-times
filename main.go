package main

import (
	"net/http"

	"github.com/Ed1123/us-visa-waiting-times/usvisa"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// // Serve static files (CSS and JS)
	// r.Static("/static", "./static")

	router.LoadHTMLFiles("./templates/table.tmpl")

	cities := usvisa.GetWaitingData()

	// Define a route to display the table
	router.GET("/", func(c *gin.Context) {
		// tmpl, err := template.ParseFiles("templates/table.tmpl")
		// if err != nil {
		// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		// 	return
		// }

		c.HTML(http.StatusOK, "table.tmpl", gin.H{"tittle": "US Visa Waiting times", "data": cities})
	})

	// Start the server
	router.Run()
}
