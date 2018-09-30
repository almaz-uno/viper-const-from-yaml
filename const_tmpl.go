package main

var constTmpl = `// {{.header}}
const (
	{{$keys := .keys}}
	{{$map := .map}}
	{{range $i, $k := $keys}} {{index $.map $k}} = "{{$k}}"
	{{end}}
)
`
