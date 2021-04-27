package templates

var HelperTemplate = `
Usage: {{.AppName}} <command> [options]

Options:
    {{.Helper}}        {{.HelperDesc}}
    {{.Table}}         {{.TableDesc}}
    {{.Name}}          {{.NameDesc}}
    {{.Migration}}         {{.MigrationDesc}}
    {{.Seeder}}            {{.SeederDesc}}

Commands:
    {{.Init}}              {{.InitDesc}}
    {{.Reset}}             {{.ResetDesc}}
    {{.Drop}}              {{.DropDesc}}
    {{.Create}}            {{.CreateDesc}}
    {{.Run}}               {{.RunDesc}}

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
var HelperCreateTemplate = `
Usage: {{.Name}} create [options]

Options:
    -t, table                           create table name for migration
    -n, name                            create file name seeder or migration
    migration                           create migration file
    seeder                              create seeder file
{{- /* end */ -}}
{{- "" }}
`
var HelperRunTemplate = `
Usage: {{.AppName}} {{.Cmd}} [options]

Options:
    {{.Migration}}              {{.MigrationDesc}}
    {{.Seeder}}                 {{.SeederDesc}}
    {{.Table}}              {{.TableDesc}}
    {{.Name}}               {{.NameDesc}}
{{- /* end */ -}}
{{- "" }}
`
