package songs

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func updateSong(ctx *gin.Context) {

	ID := ctx.PostForm("id")

	IDInInt, _ := strconv.Atoi(ID)

	title := ctx.PostForm("title")
	singer := ctx.PostForm("singer")
	writer := ctx.PostForm("writer")
	Released_date := ctx.PostForm("Released_date")

	request := songRequest{
		ID:            int(IDInInt),
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

	fmt.Println("request: ", request)

	response := songResponse{
		ID:            int(IDInInt),
		Title:         request.Title,
		Singer:        request.Singer,
		Writer:        request.Writer,
		Released_date: request.Released_date,
	}

	for i, song := range songs {
		if song.ID == int(IDInInt) {
			songs[i] = response
		}
	}

	ctx.JSON(http.StatusOK, response)

}
