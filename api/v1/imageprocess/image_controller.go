package imageprocess

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"path"
	"strings"
)

// ImageInfo 定义图片信息结构体
type ImageInfo struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// GetImages 发送轮询图
func GetImages(c *gin.Context) {
	images := []ImageInfo{} // 用于存储图片信息的切片

	dir := "resource/poll_graph" // 图片文件夹路径
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	baseURL := "http://121.36.26.12:8081/resource/poll_graph/" // 替换为你的服务器基础URL

	for _, file := range files {
		if !file.IsDir() {
			imageName := path.Base(file.Name())
			imageName = strings.TrimSuffix(imageName, path.Ext(imageName))
			imageURL := fmt.Sprintf("%s%s", baseURL, file.Name())

			// 将图片URL格式进行修改
			//imageURL = strings.Replace(imageURL, "png", "avif", -1)
			//imageURL = fmt.Sprintf("%s@%dw_%dh.avif", imageURL, 56, 56)

			imageInfo := ImageInfo{
				Name: imageName,
				URL:  imageURL,
			}
			images = append(images, imageInfo)
		}
	}

	jsonData, err := json.Marshal(images)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Data(http.StatusOK, "application/json", jsonData)
}
