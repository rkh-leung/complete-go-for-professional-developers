package migrations

import "embed"

// Go build directives that embed all sql files into executable binary
//
//go:embed *.sql
var FS embed.FS
