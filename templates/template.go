package templates

var UsageTemplate = `
Usage: {{.Name}} [command] [flag]

Options:
	-v, --version                       output the version number
	-h, --help                          output usage information
	migration                     	    create migration file
	seeder                              create seeder file
	-t, table                           make table name for migration
	-n, name                            make name file seeder or migration

Commands:
	- create, -c
	- run, -r
	- back

{{- /* end */ -}}
{{- "" }}
`
var VersionTemplate = `{{.Name}} Version {{.Version}}
{{- /* end */ -}}
{{- "" }}
`

var ErrorTemplate = `{{.Message}}: unknown command
Run 'gomig -help' for usage.
{{- /* end */ -}}
{{- "" }}
`
