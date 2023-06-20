package main

import (
	"fmt"
	"log"
	"net/http"
	"ssl_manage/core"
	"ssl_manage/server/api"

	"github.com/gin-gonic/gin"
)

var pers1 core.Config

func main() {

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/upload", func(c *gin.Context) {
		c.HTML(http.StatusOK, "upload.html", gin.H{})
	})
	router.MaxMultipartMemory = 8 << 20
	router.POST("/upload", func(c *gin.Context) {
		form, _ := c.MultipartForm()
		host := c.PostForm("host") + ":22"
		username := c.PostForm("username")
		passwd := c.PostForm("password")
		files := form.File["files"]

		//先将nginx配置文件下载到本地目录
		err := api.CopyRemoteToLocal(host, username, passwd)
		if err != nil {
			fmt.Println("忽略错误")
		}

		//读取nginx文件并修改相应的内容

		//将nginx配置文件和ssl证书文件都上传到指定目录

		for _, file := range files {
			log.Println(file.Filename)
			dst := "./uploaddir/" + file.Filename
			//上传文件至指定目录

			err := c.SaveUploadedFile(file, dst)
			if err != nil {
				fmt.Println("错误是：%s", err)
			}
			err = api.CopyFileToRemote(host, username, passwd, file.Filename)
			if err != nil {
				fmt.Println("错误是：%s", err)
				return
			}
		}
		//使用命令更新nginx配置
		pers1.SshHost = "192.168.56.210"
		pers1.SshUser = username
		pers1.SshPassword = passwd
		//pers1.SshType = "password"
		pers1.SshPort = 22
		core.NewSSH(&pers1)
		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded", len(files)))

	})
	router.Run(":8080")
}
