package templates

var UsageTemplate = `{{.Name}} Version {{.Version}}

Usage: {{.Name}} [command] [options]

Options:
	-v, --version                       output the version number
	-h, --help                          output usage information

Commands:
	- migrate
	- run
	- down

{{- /* end */ -}}
{{- "" }}
`
var VersionTemplate = `{{.Name}} Version {{.Version}}
{{- /* end */ -}}
{{- "" }}
`
