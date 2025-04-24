package ui

import "embed"

//go:embed html/pages/*.tmpl static/css/*.css static/img/* content/blog/*.md
var Files embed.FS 