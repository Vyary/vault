package ui

import (
	"embed"
	"io/fs"
)

//go:embed client/*
var WebClient embed.FS

var Web, _ = fs.Sub(WebClient, "client")
