package templates

var UsageTemplate = `
Usage: {{.Name}} <command> [options]

Options:
	-v, --version                       print version cli
	-h, --help                          output usage information
	-t, table                           create table name for migration
	-n, name                            create file name seeder or migration

Commands:
	init                                init project
	down                                delete seeder
	drop                                drop tables
	-c, create                          create migration or seeder [table name]
	-r, run                             running migration or seeder
	migration                     	    create migration file
	seeder                              create seeder file

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
