package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"os"
	"log"

	"github.com/stretchr/testify/assert"
	"github.com/gin-gonic/gin"
)

func TestPingRoute(t *testing.T) {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/albums", nil)
	router.ServeHTTP(w, req)

	dat, err := os.ReadFile("test-data/getAlbumsResult.json")
	if err != nil{
		log.Fatal(err)
	}

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, string(dat), w.Body.String())
}