package storage

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// GetFile from local storage
func GetLocalStorage(c *gin.Context) {
	year := c.Param("year")
	month := c.Param("month")
	date := c.Param("date")
	objectKey := c.Param("objectKey")
	dir, err := os.Getwd()
	if err != nil {
		log.Println(err.Error())
		return
	}
	filePath := dir + "/uploads/" + year + "/" + month + "/" + date + "/" + objectKey

	// Convert from file to byte
	file, errReadFile := os.ReadFile(filePath)
	if errReadFile != nil {
		log.Println(errReadFile.Error())
		return
	}

	c.Data(http.StatusOK, http.DetectContentType(file), file)
}
