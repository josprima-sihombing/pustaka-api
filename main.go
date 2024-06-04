package main

import (
	"net/http"

	"pustaka-api/modules/book"

	"github.com/gin-gonic/gin"
)

func main() {
	books := []book.Book{
		{
			ID:        1,
			Title:     "The Silent Ocean",
			Year:      2005,
			Publisher: "Penguin Random House",
			Rating:    8.53,
		},
		{
			ID:        2,
			Title:     "Mystery in the Shadows",
			Year:      2010,
			Publisher: "HarperCollins",
			Rating:    7.25,
		},
		{
			ID:        3,
			Title:     "Whispers of Time",
			Year:      2020,
			Publisher: "Macmillan",
			Rating:    9.12,
		},
		{
			ID:        4,
			Title:     "Journey to the Unknown",
			Year:      2018,
			Publisher: "Simon & Schuster",
			Rating:    6.78,
		},
		{
			ID:        5,
			Title:     "Echoes of Eternity",
			Year:      2003,
			Publisher: "Hachette",
			Rating:    5.94,
		},
	}

	router := gin.Default()
	router.GET("/ping", handlePingRoute)
	v1 := router.Group("/v1")

	{
		v1.GET("/books", book.HandleGetBooks(&books))
		v1.POST("/books", book.HandleCreateBook(&books))
	}

	router.Run()
}

func handlePingRoute(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
