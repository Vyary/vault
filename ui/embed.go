package ui

import (
	"embed"
	"io/fs"
)

//go:embed static/*
var WebClient embed.FS

var Web, _ = fs.Sub(WebClient, "static")

//go:embed index.html
var Index embed.FS
