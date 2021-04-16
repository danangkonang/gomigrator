package templates

var HelperTemplate = `
Usage: {{.Name}} <command> [options]

Options:
    -v, --version                       print version cli
    -h, --help                          output usage information
    -t, table                           create table name for migration
    -n, name                            create file name seeder or migration
    -m, migration                       create migration file
    -s, seeder                          create seeder file

Commands:
    init                                init project
    down                                delete seeder
    drop                                drop tables
    create                              create migration or seeder [table name]
    run                                 running migration or seeder

{{- /* end */ -}}
{{- "" }}
`
var VersionTemplate = `{{.Name}} version {{.Version}}
{{- /* end */ -}}
{{- "" }}
`

var ErrorTemplate = `{{.Message}}: unknown command
Run 'gomig -help' for usage.
{{- /* end */ -}}
{{- "" }}
`
