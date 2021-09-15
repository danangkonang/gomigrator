package templates

import (
	"fmt"
	"html/template"
	"os"
)

type ComandUsage struct {
	CmdName  string
	CmdAlias string
	CmdDesc  string
}

type FlagCmd struct {
	FlagName  string
	FlagAlias string
	FlagDesc  string
}

type Helper struct {
	Usage    string
	Version  string
	Error    string
	Option   []*ComandUsage
	Argument []*FlagCmd
}

var ErrorTmp = `{{ .Error }}

see 'gomigator --help'
`

var VersionTmp = `Version: {{ .Version }}
`

var HelperTmp = `
Usage: {{ .Usage }}
{{ if .Option }}
Commands:
{{- range .Option}}
  {{ .CmdName }}    {{"\t"}}{{ .CmdDesc }}{{ end -}}
{{ end }}

Options:
{{- range .Argument}}
  {{ .FlagName }}        {{"\t"}}{{ .FlagDesc }}{{ end }}

`

func PrintTemplate(temp string, data interface{}) {
	tmpl, err := template.New("helper").Parse(temp)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	os.Exit(0)
}
