package storage

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// SaveFile in local storage
func SaveLocalStorage(c *gin.Context) {
	addr := getEnv("URL_STORAGE", "https://baseApi.id/")

	defer c.Request.MultipartForm.RemoveAll()

	file, _ := c.FormFile("file")
	body, err := file.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File must be provided"})
		return
	}
	defer body.Close()

	currentTime := time.Now()
	datePath := currentTime.Format("2006/01/02")
	objectKeyPrefix := strconv.FormatInt(time.Now().Unix(), 10)
	filename := objectKeyPrefix + "-" + url.QueryEscape(file.Filename)

	dir, err := os.Getwd()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	usedPath := "/uploads/" + datePath
	errMkdir := os.MkdirAll(dir+"/"+usedPath, os.ModePerm)
	if errMkdir != nil {
		log.Println("errmkdir: " + errMkdir.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": errMkdir.Error()})
		return
	}

	fileLocation := filepath.Join(dir, usedPath, filename)
	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	defer targetFile.Close()
	if _, err := io.Copy(targetFile, body); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	urlPath := "/get/" + datePath
	c.JSON(http.StatusOK, gin.H{
		"data": map[string]string{"filename": addr + urlPath + "/" + filename},
	})
}
