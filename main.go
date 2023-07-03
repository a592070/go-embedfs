package main

import (
	"embed"
	"github.com/gin-gonic/gin"
	"io/fs"
	"net/http"
)

//go:embed public/*
var app embed.FS

func main() {
	engine := gin.Default()

	subFS, err := fs.Sub(app, "public")
	if err != nil {
		panic(err)
	}
	engine.StaticFS("/", http.FS(subFS))

	engine.Run(":3000")
}
