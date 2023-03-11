package assets

import "github.com/gobuffalo/packr/v2"

//go:generate packr2 clean
//go:generate packr2
//go:generate gofmt -w .

//Assets - assets defines packr box of csv folder
var Assets = packr.New("assets", "../../../csv")
