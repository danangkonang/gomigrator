package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"text/template"

	"github.com/danangkonang/migration-go-cli/app/command"
	"github.com/danangkonang/migration-go-cli/app/model"
)

var (
	help    bool
	version bool
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
	Option   []*ComandUsage
	Argument []*FlagCmd
}

var helperTmp = `
Usage: create {{"\n"}}
{{- if .Option }}
Commands:
{{- range .Option}}
  {{ .CmdName }}    {{"\t"}}{{ .CmdDesc }}{{ end -}}
{{"\n"}}
{{ end }}
Options:
{{- range .Argument}}
  {{ .FlagName }}  {{"\t"}}{{ .FlagDesc }}{{ end -}}
{{"\n"}}
`

func printHeler(temp string, data interface{}) {
	tmpl, err := template.New("heler").Parse(temp)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}

func main() {
	flag.BoolVar(&help, "h", false, "help")
	flag.BoolVar(&help, "help", false, "help")
	flag.BoolVar(&version, "v", false, "version")
	flag.BoolVar(&version, "version", false, "version")
	flag.Parse()
	if help || len(os.Args[1:]) == 0 {
		fmt.Println("show help 1")
		os.Exit(0)
	}
	if version {
		fmt.Println("show version")
		os.Exit(0)
	}
	var i model.Init
	init := flag.NewFlagSet("init", flag.ExitOnError)
	init.Func("driver", "set up db driver", func(s string) error {
		i.Driver = s
		return nil
	})
	init.Func("port", "set up db port", func(s string) error {
		port, _ := strconv.Atoi(s)
		i.Port = port
		return nil
	})
	init.Func("host", "set up db host", func(s string) error {
		i.Host = s
		return nil
	})
	init.Func("database", "set up db host", func(s string) error {
		i.DAtabase = s
		return nil
	})
	init.Func("user", "set up db user", func(s string) error {
		i.User = s
		return nil
	})
	init.Func("password", "set up db password", func(s string) error {
		i.Password = s
		return nil
	})

	create := flag.NewFlagSet("create", flag.ExitOnError)
	var c model.Create
	migration := flag.NewFlagSet("migration", flag.ContinueOnError)
	migration.Func("name", "generate file name", func(s string) error {
		c.FileName = s
		return nil
	})
	migration.Func("table", "generate file name", func(s string) error {
		c.TableName = s
		return nil
	})
	seeder := flag.NewFlagSet("seeder", flag.ContinueOnError)
	seeder.Func("name", "generate file name", func(s string) error {
		c.FileName = s
		return nil
	})
	seeder.Func("table", "generate file name", func(s string) error {
		c.TableName = s
		return nil
	})

	up := flag.NewFlagSet("up", flag.ExitOnError)
	down := flag.NewFlagSet("down", flag.ExitOnError)
	reset := flag.NewFlagSet("reset", flag.ExitOnError)

	if len(os.Args) < 2 {
		fmt.Println("helper")
		os.Exit(0)
	}
	switch os.Args[1] {
	case "init":
		order := []string{"host", "port", "database", "driver", "user", "password"}
		init.Usage = func() {
			var cmdFlag []*FlagCmd
			for _, item := range order {
				flg := init.Lookup(item)
				cmdFlag = append(cmdFlag, &FlagCmd{
					FlagName: fmt.Sprintf("--%s", flg.Name),
					FlagDesc: flg.Usage,
				})
			}
			hlp := &Helper{
				Argument: cmdFlag,
			}
			printHeler(helperTmp, hlp)
		}
		init.Parse(os.Args[2:])
		// cli.Init(&i)
		command.Init(&i)
	case "create":
		create.Usage = func() {
			var cmdOption []*ComandUsage
			cmdOption = append(cmdOption, &ComandUsage{
				CmdName: migration.Name(),
				CmdDesc: "migration file",
			})
			cmdOption = append(cmdOption, &ComandUsage{
				CmdName: seeder.Name(),
				CmdDesc: "migration file",
			})
			var cmdFlag []*FlagCmd
			order := []string{"table", "name"}
			for _, item := range order {
				flg := migration.Lookup(item)
				cmdFlag = append(cmdFlag, &FlagCmd{
					FlagName: fmt.Sprintf("--%s", flg.Name),
					FlagDesc: flg.Usage,
				})
			}
			hlp := &Helper{
				Option:   cmdOption,
				Argument: cmdFlag,
			}
			printHeler(helperTmp, hlp)
		}
		create.Parse(os.Args[2:])
		handleCreate(os.Args, migration, seeder, &c)
	case "up":
		up.Parse(os.Args[2:])
		// HandleUp(os.Args)
	case "down":
		down.Parse(os.Args[2:])
		// HandleDown(os.Args)
	case "reset":
		reset.Parse(os.Args[2:])
	default:
		fmt.Println("helper")
		os.Exit(0)
	}
}

func handleCreate(argument []string, migration, seeder *flag.FlagSet, c *model.Create) {
	switch argument[2] {
	case "migration":
		migration.Usage = func() {
			var cmdFlag []*FlagCmd
			order := []string{"table", "name"}
			for _, item := range order {
				flg := migration.Lookup(item)
				cmdFlag = append(cmdFlag, &FlagCmd{
					FlagName: fmt.Sprintf("--%s", flg.Name),
					FlagDesc: flg.Usage,
				})
			}
			hlp := &Helper{
				Argument: cmdFlag,
			}
			printHeler(helperTmp, hlp)
		}
		migration.Parse(argument[3:])
		command.CreateMigtaion(c)
	case "seeder":
		seeder.Parse(argument[3:])
		// cli.CreateSeeder(&c)
	default:
		fmt.Println("helper")
		os.Exit(0)
	}
}

// func main() {
// 	arrCmd := os.Args[1:]
// 	if len(arrCmd) == 0 {
// 		helper.PrintHelper()
// 		return
// 	}
// 	runCmd()
// }

// func runCmd() {
// 	usrCmd := os.Args[1]
// 	switch usrCmd {
// 	/*
// 	* print helper cli
// 	 */
// 	case "-h":
// 		helper.PrintHelper()
// 	case "--help":
// 		helper.PrintHelper()

// 	/*
// 	* print version aplication cli
// 	 */
// 	case "-v":
// 		helper.PrintVersion()
// 	case "--version":
// 		helper.PrintVersion()

// 	/*
// 	* comand for init db migration and seeder directory
// 	* expanple: go run main.go init
// 	 */
// 	case "init":
// 		command.InitialV2()

// 	/*
// 	* comand for generate migration or seeder file
// 	* expanple: go run main.go create migration --name <table name>
// 	* file will be save to db/migration/<filename>
// 	* expanple: go run main.go create seeder --name <file name> --table <table name>
// 	* file will be save to db/seeder/<filename>
// 	 */
// 	case "create":
// 		command.MigrationSeederCreate()
// 	case "-c":
// 		command.MigrationSeederCreate()

// 	/*
// 	* comand for running migration or seeder
// 	* expanple: go run main.go run migration
// 	* expanple: go run main.go run seeder
// 	 */
// 	case "run":
// 		command.MigrationSeederRun()
// 	case "-r":
// 		command.MigrationSeederRun()

// 	/*
// 	* comand for delete all data in table
// 	* expanple: go run main.go reset
// 	 */
// 	case "reset":
// 		command.EmtySeederData()

// 	/*
// 	* comand for delete all tables in database
// 	* expanple: go run main.go drop
// 	 */
// 	case "drop":
// 		command.DropTableRun()
// 	default:
// 		helper.PrintHelper()
// 	}
// }
