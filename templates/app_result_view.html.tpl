{{ define `root`}}
{{ if .IsSource -}}
```{{ .SourceCodeType }}
{{ end -}}
{{ .MainContent }}
{{ if .IsSource -}}
```
{{ end -}}
{{end}}
