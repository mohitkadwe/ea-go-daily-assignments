package songs

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getSong(ctx *gin.Context) {

	id := ctx.Query("id")

	IdInInt, err := strconv.Atoi(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	song := songs[IdInInt]

	releaseDate := song.Released_date

	response := songResponse{
		ID:            song.ID,
		Title:         song.Title,
		Singer:        song.Singer,
		Writer:        song.Writer,
		Released_date: releaseDate,
	}

	fmt.Println(response)

	ctx.JSON(http.StatusOK, response)

}
