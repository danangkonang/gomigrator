package templates

var UsageTemplate = `
Usage: {{.Name}} [command] [flag]

Options:
	-v, --version                       output the version number
	-h, --help                          output usage information
	-m, --migration                     create migration file
	-s, --seeder                        create seeder file

Commands:
	- create
	- run
	- back

{{- /* end */ -}}
{{- "" }}
`
var VersionTemplate = `{{.Name}} Version {{.Version}}
{{- /* end */ -}}
{{- "" }}
`
