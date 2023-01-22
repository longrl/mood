package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/longrl/mood/response"
	"strings"
)

func Upload(c *gin.Context) {
	file, _ := c.FormFile("file")
	var name string
	path := "C:/Users/uu/Desktop/mood/server/static"
	ext := getFileExt(file.Filename)
	if ext == "mp3" {
		name = "mp3"
	} else {
		name = "image"
	}
	dst := path + "/" + name + "/" + file.Filename
	// 上传文件至指定的完整文件路径
	c.SaveUploadedFile(file, dst)

	response.JSON(c, gin.H{
		"url": "http://127.0.0.1:5500/" + name + "/" + file.Filename,
	})
}

func getFileExt(name string) string {
	index := strings.LastIndex(name, ".")
	if index == len(name)-1 {
		return ""
	}
	return name[index+1:]
}
