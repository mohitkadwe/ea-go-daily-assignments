package songs

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func makeACallForSong(reqBody songRequest) *httptest.ResponseRecorder {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/songs", getSong)

	ID := strconv.Itoa(reqBody.ID)

	request, _ := http.NewRequest("GET", "/songs?id="+ID, nil)
	response := httptest.NewRecorder()

	r.ServeHTTP(response, request)

	return response

}

func TestGetSongByQuery(t *testing.T) {

	reqBody := songRequest{
		ID: 1,
	}

	response := makeACallForSong(reqBody)
	assert.Equal(t, http.StatusOK, response.Code)

	responseBody, _ := io.ReadAll(response.Body)

	res := songResponse{}

	json.Unmarshal(responseBody, &res)

	assert.Equal(t, 2, res.ID)
	assert.Equal(t, "Song 2", res.Title)
	assert.Equal(t, "Singer 2", res.Singer)
	assert.Equal(t, "Writer 2", res.Writer)
	assert.Equal(t, "2021-01-02", res.Released_date)

}

func makeACallToAddSong(reqBody songRequest) *httptest.ResponseRecorder {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.POST("/songs", createSong)

	request, _ := http.NewRequest("POST", "/songs", nil)
	request.PostForm = map[string][]string{
		"title":         {reqBody.Title},
		"singer":        {reqBody.Singer},
		"writer":        {reqBody.Writer},
		"Released_date": {reqBody.Released_date},
	}
	response := httptest.NewRecorder()

	r.ServeHTTP(response, request)

	return response

}

func TestCreateSong(t *testing.T) {

	reqBody := songRequest{
		Title:         "Song 3",
		Singer:        "Singer 3",
		Writer:        "Writer 3",
		Released_date: "2021-01-03",
	}

	response := makeACallToAddSong(reqBody)
	assert.Equal(t, http.StatusOK, response.Code)
	responseBody, _ := io.ReadAll(response.Body)

	res := songResponse{}

	json.Unmarshal(responseBody, &res)

	assert.Equal(t, 3, res.ID)
	assert.Equal(t, "Song 3", res.Title)
	assert.Equal(t, "Singer 3", res.Singer)
	assert.Equal(t, "Writer 3", res.Writer)
	assert.Equal(t, "2021-01-03", res.Released_date)
}

func makeACallToUpdateSong(reqBody songRequest) *httptest.ResponseRecorder {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.PUT("/songs", updateSong)

	ID := strconv.Itoa(reqBody.ID)

	request, _ := http.NewRequest("PUT", "/songs", nil)
	request.PostForm = map[string][]string{
		"id":            {ID},
		"title":         {reqBody.Title},
		"singer":        {reqBody.Singer},
		"writer":        {reqBody.Writer},
		"Released_date": {reqBody.Released_date},
	}
	response := httptest.NewRecorder()

	r.ServeHTTP(response, request)

	return response

}

func TestUpdateSong(t *testing.T) {

	reqBody := songRequest{
		ID:            1,
		Title:         "Song 1 new",
		Singer:        "Singer 1 new",
		Writer:        "Writer 1 new",
		Released_date: "2021-01-01",
	}

	response := makeACallToUpdateSong(reqBody)
	assert.Equal(t, http.StatusOK, response.Code)
	responseBody, _ := io.ReadAll(response.Body)

	res := songResponse{}

	json.Unmarshal(responseBody, &res)

	assert.Equal(t, 1, res.ID)
	assert.Equal(t, "Song 1 new", res.Title)
	assert.Equal(t, "Singer 1 new", res.Singer)
	assert.Equal(t, "Writer 1 new", res.Writer)
	assert.Equal(t, "2021-01-01", res.Released_date)
}

func makeACallToDeleteSong(reqBody songRequest) *httptest.ResponseRecorder {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.DELETE("/songs", deleteSong)

	ID := strconv.Itoa(reqBody.ID)

	request, _ := http.NewRequest("DELETE", "/songs?id="+ID, nil)
	response := httptest.NewRecorder()

	r.ServeHTTP(response, request)

	return response

}

func TestDeleteSong(t *testing.T) {

	reqBody := songRequest{
		ID: 1,
	}

	response := makeACallToDeleteSong(reqBody)
	assert.Equal(t, http.StatusOK, response.Code)
	responseBody, _ := io.ReadAll(response.Body)

	res := songResponse{}

	json.Unmarshal(responseBody, &res)

	assert.Equal(t, reqBody.ID, res.ID)

}
