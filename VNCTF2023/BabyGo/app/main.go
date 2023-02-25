package main

import (
	"encoding/gob"
	"fmt"
	"github.com/PaulXu-cn/goeval"
	"github.com/duke-git/lancet/cryptor"
	"github.com/duke-git/lancet/fileutil"
	"github.com/duke-git/lancet/random"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type User struct {
	Name  string
	Path  string
	Power string
}

func main() {
	r := gin.Default()
	store := cookie.NewStore(random.RandBytes(16))
	r.Use(sessions.Sessions("session", store))
	r.LoadHTMLGlob("template/*")

	r.GET("/", func(c *gin.Context) {
		userDir := "/tmp/" + cryptor.Md5String(c.ClientIP()+"VNCTF2023GoGoGo~") + "/"
		session := sessions.Default(c)
		session.Set("shallow", userDir)
		session.Save()
		fileutil.CreateDir(userDir)
		gobFile, _ := os.Create(userDir + "user.gob")
		user := User{Name: "ctfer", Path: userDir, Power: "low"}
		encoder := gob.NewEncoder(gobFile)
		encoder.Encode(user)
		if fileutil.IsExist(userDir) && fileutil.IsExist(userDir+"user.gob") {
			c.HTML(200, "index.html", gin.H{"message": "Your path: " + userDir})
			return
		}
		c.HTML(500, "index.html", gin.H{"message": "failed to make user dir"})
	})

	r.GET("/upload", func(c *gin.Context) {
		c.HTML(200, "upload.html", gin.H{"message": "upload me!"})
	})

	r.POST("/upload", func(c *gin.Context) {
		session := sessions.Default(c)
		if session.Get("shallow") == nil {
			c.Redirect(http.StatusFound, "/")
		}
		userUploadDir := session.Get("shallow").(string) + "uploads/"
		fileutil.CreateDir(userUploadDir)
		file, err := c.FormFile("file")
		if err != nil {
			c.HTML(500, "upload.html", gin.H{"message": "no file upload"})
			return
		}
		ext := file.Filename[strings.LastIndex(file.Filename, "."):]
		if ext == ".gob" || ext == ".go" {
			c.HTML(500, "upload.html", gin.H{"message": "Hacker!"})
			return
		}
		filename := userUploadDir + file.Filename
		if fileutil.IsExist(filename) {
			fileutil.RemoveFile(filename)
		}
		err = c.SaveUploadedFile(file, filename)
		if err != nil {
			c.HTML(500, "upload.html", gin.H{"message": "failed to save file"})
			return
		}
		c.HTML(200, "upload.html", gin.H{"message": "file saved to " + filename})
	})

	r.GET("/unzip", func(c *gin.Context) {
		session := sessions.Default(c)
		if session.Get("shallow") == nil {
			c.Redirect(http.StatusFound, "/")
		}
		userUploadDir := session.Get("shallow").(string) + "uploads/"
		files, _ := fileutil.ListFileNames(userUploadDir)
		destPath := filepath.Clean(userUploadDir + c.Query("path"))
		for _, file := range files {
			if fileutil.MiMeType(userUploadDir+file) == "application/zip" {
				err := fileutil.UnZip(userUploadDir+file, destPath)
				if err != nil {
					c.HTML(200, "zip.html", gin.H{"message": "failed to unzip file"})
					return
				}
				fileutil.RemoveFile(userUploadDir + file)
			}
		}
		c.HTML(200, "zip.html", gin.H{"message": "success unzip"})
	})

	r.GET("/backdoor", func(c *gin.Context) {
		session := sessions.Default(c)
		if session.Get("shallow") == nil {
			c.Redirect(http.StatusFound, "/")
		}
		userDir := session.Get("shallow").(string)
		if fileutil.IsExist(userDir + "user.gob") {
			file, _ := os.Open(userDir + "user.gob")
			decoder := gob.NewDecoder(file)
			var ctfer User
			decoder.Decode(&ctfer)
			if ctfer.Power == "admin" {
				eval, err := goeval.Eval("", "fmt.Println(\"Good\")", c.DefaultQuery("pkg", "fmt"))
				if err != nil {
					fmt.Println(err)
				}
				c.HTML(200, "backdoor.html", gin.H{"message": string(eval)})
				return
			} else {
				c.HTML(200, "backdoor.html", gin.H{"message": "low power"})
				return
			}
		} else {
			c.HTML(500, "backdoor.html", gin.H{"message": "no such user gob"})
			return
		}
	})

	r.Run(":8080")
}
