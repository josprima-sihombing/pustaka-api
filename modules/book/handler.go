package book

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleGetBooks(books *[]Book) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"data": books,
		})
	}
}

func HandleCreateBook(books *[]Book) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var bookInput CreateBookInput

		err := ctx.ShouldBindJSON(&bookInput)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid JSON request",
			})

			return
		}

		newBook := Book{
			ID:        2,
			Title:     bookInput.Title,
			Year:      bookInput.Year,
			Publisher: bookInput.Publisher,
			Rating:    bookInput.Rating,
		}

		*books = append(*books, newBook)

		ctx.JSON(http.StatusOK, gin.H{
			"data": books,
		})
	}
}
