package static

import (
	"embed"
)

//go:embed css/* js/* images/*
var StaticFiles embed.FS

//go:embed templates/*
var Templates embed.FS 