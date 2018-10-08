package main

var constTmpl = `package {{.pkg}}

// {{.header}}
const (
	{{$keys := .keys}}
	{{$map := .map}}
	{{$prefix := .prefix}}
	{{range $i, $k := $keys}}{{$prefix}}{{index $.map $k}} = "{{$k}}"
	{{end}}
)`
