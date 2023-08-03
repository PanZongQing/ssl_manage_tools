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
	// var host, port string
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "upload.html", gin.H{})
	})
	router.MaxMultipartMemory = 8 << 20
	router.POST("/", func(c *gin.Context) {
		form, _ := c.MultipartForm()
		Auto_domain := c.PostForm("domain")
		manual_domian := c.PostForm("manual_domain")
		port := c.PostForm("port")
		host := c.PostForm("host")
		username := c.PostForm("username")
		passwd := c.PostForm("password")
		files := form.File["files"]
		fmt.Println(Auto_domain)

		//判断是用户填写的文件名还是选择的域名
		if manual_domian == "" {
			switch Auto_domain {
			case "www.polixir.site":
				host = "192.168.56.210"
				port = "22"
			case "www.polixir.ai":
				host = "192.168.56.210"
				port = "22"
			case "www.agit.ai":
				host = "192.168.56.210"
				port = "22"
			case "service.agit.site":
				host = "106.75.182.54"
				port = "122"
			default:
				fmt.Println("就这样了")

			}
		} else {
			fmt.Println("程序出错了")
			return
		}
		fmt.Println(host, port)

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
		pers1.SshHost = host
		pers1.SshUser = username
		pers1.SshPassword = passwd
		pers1.SshType = "password"
		pers1.SshPort = port
		core.NewSSH(&pers1)

		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded", len(files)))

	})
	router.Run(":8080")
}
