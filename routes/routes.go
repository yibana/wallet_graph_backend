package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"wallet_graph_backend/OpenApi"
)

func Readme(c *gin.Context) {
	filePath := filepath.Join(".", "README.md")
	markdownContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Error reading file: %s", err.Error()))
		return
	}
	// 渲染模板，并将markdownContent变量替换为动态的Markdown内容
	c.HTML(200, "markdown.html", gin.H{
		"MarkdownContent": string(markdownContent),
		"title":           "README.md",
	})
}

func OpenApiMistTrack(c *gin.Context) {
	methodValue := c.Param("method")
	coin := c.Query("coin")
	address := c.Query("address")
	var rsp OpenApi.OpenApiResult
	var err error
	defer func() {
		if err != nil {
			rsp.Status = "error"
			rsp.Error = err.Error()
		} else {
			rsp.Status = "ok"
		}
		c.JSON(http.StatusOK, rsp)
	}()
	rsp, err = OpenApi.OpenApiMistTrack(methodValue, coin, address)
}
