package songs

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func createSong(ctx *gin.Context) {

	title := ctx.PostForm("title")
	singer := ctx.PostForm("singer")
	writer := ctx.PostForm("writer")
	Released_date := ctx.PostForm("Released_date")

	request := songRequest{
		Title:         title,
		Singer:        singer,
		Writer:        writer,
		Released_date: Released_date,
	}
	err := ctx.Bind(&request)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := songResponse{
		ID:            len(songs) + 1,
		Title:         request.Title,
		Singer:        request.Singer,
		Writer:        request.Writer,
		Released_date: request.Released_date,
	}

	result := append(songs, songResponse{
		ID:            len(songs) + 1,
		Title:         request.Title,
		Singer:        request.Singer,
		Writer:        request.Writer,
		Released_date: request.Released_date,
	})

	songs = result

	ctx.JSON(http.StatusOK, response)

}
