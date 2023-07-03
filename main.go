package main

import (
	"embed"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"io/fs"
	"net/http"
)

var (
	//go:embed client/build
	react embed.FS

	////go:embed public
	//app embed.FS
)

func main() {
	engine := gin.Default()

	engine.Use(EmbedReact("/", "client/build", react))

	engine.Run(":3001")
}

func EmbedReact(urlPrefix, buildDirectory string, em embed.FS) gin.HandlerFunc {
	dir := static.LocalFile(buildDirectory, true)
	embedDir, _ := fs.Sub(em, buildDirectory)
	fileserver := http.FileServer(http.FS(embedDir))

	if urlPrefix != "" {
		fileserver = http.StripPrefix(urlPrefix, fileserver)
	}

	return func(c *gin.Context) {
		if !dir.Exists(urlPrefix, c.Request.URL.Path) {
			c.Request.URL.Path = "/"
		}
		fileserver.ServeHTTP(c.Writer, c.Request)
		c.Abort()
	}
}
