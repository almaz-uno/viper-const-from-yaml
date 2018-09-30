package main

var constTmpl = `package {{.pkg}}

// {{.header}}
const (
	{{$keys := .keys}}
	{{$map := .map}}
	{{range $i, $k := $keys}}{{index $.map $k}} = "{{$k}}"
	{{end}}
)`
