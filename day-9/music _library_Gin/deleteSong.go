package songs

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func deleteSong(ctx *gin.Context) {

	id := ctx.Query("id")

	IdInInt, err := strconv.Atoi(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	for i, song := range songs {
		if song.ID == IdInInt {
			songs = append(songs[:i], songs[i+1:]...)
		}
	}

	fmt.Println("songs: ", songs)

	ctx.JSON(http.StatusOK, gin.H{
		"ID": IdInInt,
	})

}
