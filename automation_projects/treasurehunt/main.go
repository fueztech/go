package main

import (
	"net/http"
	"treasurehunt/treasures"

	"github.com/gin-gonic/gin"
)

// homePage this is the landing page for the website
func homePage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

// getScriptures responds with the list of all scriptures as JSON
func getScriptures(c *gin.Context) {
	weeklyScripture := treasures.ScriptureCollection()
	c.IndentedJSON(http.StatusOK, weeklyScripture) // Also will add the location
}

func postScriptures(c *gin.Context) {
	var addTreasure treasures.SpiritualTreasure

	if err := c.BindJSON(&addTreasure); err != nil {
		return
	}
	// Add the new set of scriptures to the slice of the scriptures slice
	// In treasures module.
	treasures.AddScriptures(addTreasure)

	//Scriptures = append(Scriptures, newTreasure)
	c.IndentedJSON(http.StatusCreated, addTreasure)
}

func main() {
	router := gin.Default()
	router.GET("/home", homePage)
	router.LoadHTMLGlob("templates/*.html")
	router.GET("/home/scriptures", getScriptures)
	router.POST("/home/scriptures", postScriptures)
	router.Run("localhost:8080")
}
