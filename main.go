package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/template"

	"github.com/danangkonang/gomigrator/app/command"
	"github.com/danangkonang/gomigrator/app/model"
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
	Usage    string
	Version  string
	Error    string
	Option   []*ComandUsage
	Argument []*FlagCmd
}

var versionTmp = `Version: {{ .Version }}
`

var errorTmp = `unknow comand '{{ .Error }}'

see 'gomigator --help'
`

var helperTmp = `
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

func printTemplate(temp string, data interface{}) {
	tmpl, err := template.New("heler").Parse(temp)
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

func globalHelp() {
	hlp := &Helper{
		Usage: "gomigator [COMAND] [OPTIONS]",
		Option: []*ComandUsage{
			{
				CmdName: "init",
				CmdDesc: "generate db directory for",
			},
			{
				CmdName: "create",
				CmdDesc: "create migration or seeder file",
			},
			{
				CmdName: "up",
				CmdDesc: "exect migration to database",
			},
			{
				CmdName: "down",
				CmdDesc: "drop migration on databse",
			},
			{
				CmdName: "migration",
				CmdDesc: "generate type migration",
			},
			{
				CmdName: "seeder",
				CmdDesc: "generate type seeder",
			},
		},
		Argument: []*FlagCmd{
			{
				FlagName: "-h, --help",
				FlagDesc: "print help gomigrator",
			},
			{
				FlagName: "-v, --version",
				FlagDesc: "print version gomigrator",
			},
			{
				FlagName: "--table",
				FlagDesc: "table name",
			},
			{
				FlagName: "--tables",
				FlagDesc: "list tables",
			},
			{
				FlagName: "--name",
				FlagDesc: "generate file name",
			},
		},
	}
	printTemplate(helperTmp, hlp)
}

func main() {
	flag.BoolVar(&help, "h", false, "help")
	flag.BoolVar(&help, "help", false, "help")
	flag.BoolVar(&version, "v", false, "version")
	flag.BoolVar(&version, "version", false, "version")
	flag.Parse()
	if help || len(os.Args[1:]) == 0 {
		globalHelp()
	}
	if version {
		hlp := &Helper{
			Version: "0.0.8",
		}
		printTemplate(versionTmp, hlp)
	}
	var i model.Init

	/*
	*
	 */
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

	/*
	*
	 */
	create := flag.NewFlagSet("create", flag.ExitOnError)
	var c model.Create
	migration := flag.NewFlagSet("migration", flag.ContinueOnError)
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

	/*
	*
	 */
	up := flag.NewFlagSet("up", flag.ExitOnError)
	var upM model.UpDown
	upMigration := flag.NewFlagSet("migration", flag.ContinueOnError)
	upMigration.Func("tables", "list file name", func(s string) error {
		upM.Tables = strings.Fields(s)
		return nil
	})
	upSeeder := flag.NewFlagSet("seeder", flag.ContinueOnError)
	upSeeder.Func("tables", "list file name", func(s string) error {
		upM.Tables = strings.Fields(s)
		return nil
	})

	/*
	*
	 */
	down := flag.NewFlagSet("down", flag.ExitOnError)
	var dwM model.UpDown
	// remove table migartion
	downMigration := flag.NewFlagSet("migration", flag.ContinueOnError)
	downMigration.Func("tables", "list file name", func(s string) error {
		dwM.Tables = strings.Fields(s)
		return nil
	})

	// remove seeder data
	downSeeder := flag.NewFlagSet("seeder", flag.ContinueOnError)
	downSeeder.Func("tables", "list file name", func(s string) error {
		dwM.Tables = strings.Fields(s)
		return nil
	})

	if len(os.Args) < 2 {
		globalHelp()
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
				Usage:    "gomigator init [OPTIONS]",
				Argument: cmdFlag,
			}
			printTemplate(helperTmp, hlp)
		}
		init.Parse(os.Args[2:])
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
			order := []string{"table"}
			for _, item := range order {
				flg := migration.Lookup(item)
				cmdFlag = append(cmdFlag, &FlagCmd{
					FlagName: fmt.Sprintf("--%s", flg.Name),
					FlagDesc: flg.Usage,
				})
			}
			hlp := &Helper{
				Usage:    "gomigator create [COMMAND] [OPTIONS]",
				Option:   cmdOption,
				Argument: cmdFlag,
			}
			printTemplate(helperTmp, hlp)
		}
		create.Parse(os.Args[2:])
		createHandle(os.Args, migration, seeder, &c)
	case "up":
		up.Usage = func() {
			var cmdOption []*ComandUsage
			cmdOption = append(cmdOption, &ComandUsage{
				CmdName: upMigration.Name(),
				CmdDesc: "migration file",
			})
			cmdOption = append(cmdOption, &ComandUsage{
				CmdName: upSeeder.Name(),
				CmdDesc: "migration file",
			})
			var cmdFlag []*FlagCmd
			order := []string{"tables"}
			for _, item := range order {
				flg := upMigration.Lookup(item)
				cmdFlag = append(cmdFlag, &FlagCmd{
					FlagName: fmt.Sprintf("--%s", flg.Name),
					FlagDesc: flg.Usage,
				})
			}
			hlp := &Helper{
				Usage:    "gomigator up [COMMAND] [OPTIONS]",
				Option:   cmdOption,
				Argument: cmdFlag,
			}
			printTemplate(helperTmp, hlp)
		}
		up.Parse(os.Args[2:])
		upHandle(os.Args, upMigration, upSeeder, &upM)
	case "down":
		down.Usage = func() {
			var cmdOption []*ComandUsage
			cmdOption = append(cmdOption, &ComandUsage{
				CmdName: upMigration.Name(),
				CmdDesc: "migration file",
			})
			cmdOption = append(cmdOption, &ComandUsage{
				CmdName: upSeeder.Name(),
				CmdDesc: "migration file",
			})
			var cmdFlag []*FlagCmd
			order := []string{"tables"}
			for _, item := range order {
				flg := upMigration.Lookup(item)
				cmdFlag = append(cmdFlag, &FlagCmd{
					FlagName: fmt.Sprintf("--%s", flg.Name),
					FlagDesc: flg.Usage,
				})
			}
			hlp := &Helper{
				Usage:    "gomigator down [COMMAND] [OPTIONS]",
				Option:   cmdOption,
				Argument: cmdFlag,
			}
			printTemplate(helperTmp, hlp)
		}
		down.Parse(os.Args[2:])
		downHandle(os.Args, downMigration, downSeeder, &dwM)
	default:
		hlp := &Helper{
			Error: os.Args[1],
		}
		printTemplate(errorTmp, hlp)
	}
}

func downHandle(argument []string, downMigration, downSeeder *flag.FlagSet, c *model.UpDown) {
	switch argument[2] {
	case "migration":
		downMigration.Usage = func() {
			var cmdFlag []*FlagCmd
			order := []string{"tables"}
			for _, item := range order {
				flg := downMigration.Lookup(item)
				cmdFlag = append(cmdFlag, &FlagCmd{
					FlagName: fmt.Sprintf("--%s", flg.Name),
					FlagDesc: flg.Usage,
				})
			}
			hlp := &Helper{
				Usage:    "gomigator down migration [OPTIONS]",
				Argument: cmdFlag,
			}
			printTemplate(helperTmp, hlp)
		}
		downMigration.Parse(argument[3:])
		command.DownMigration(c)
	case "seeder":
		downSeeder.Usage = func() {
			var cmdFlag []*FlagCmd
			order := []string{"tables"}
			for _, item := range order {
				flg := downSeeder.Lookup(item)
				cmdFlag = append(cmdFlag, &FlagCmd{
					FlagName: fmt.Sprintf("--%s", flg.Name),
					FlagDesc: flg.Usage,
				})
			}
			hlp := &Helper{
				Usage:    "gomigator down seeder [OPTIONS]",
				Argument: cmdFlag,
			}
			printTemplate(helperTmp, hlp)
		}
		downSeeder.Parse(argument[3:])
		command.DownSeeder(c)
	default:
		hlp := &Helper{
			Error: os.Args[2],
		}
		printTemplate(errorTmp, hlp)
	}
}

func upHandle(argument []string, upMigration, upSeeder *flag.FlagSet, c *model.UpDown) {
	switch argument[2] {
	case "migration":
		upMigration.Usage = func() {
			var cmdFlag []*FlagCmd
			order := []string{"tables"}
			for _, item := range order {
				flg := upMigration.Lookup(item)
				cmdFlag = append(cmdFlag, &FlagCmd{
					FlagName: fmt.Sprintf("--%s", flg.Name),
					FlagDesc: flg.Usage,
				})
			}
			hlp := &Helper{
				Usage:    "gomigator up migration [OPTIONS]",
				Argument: cmdFlag,
			}
			printTemplate(helperTmp, hlp)
		}
		upMigration.Parse(argument[3:])
		command.UpMigration(c)
	case "seeder":
		upSeeder.Usage = func() {
			var cmdFlag []*FlagCmd
			order := []string{"tables"}
			for _, item := range order {
				flg := upSeeder.Lookup(item)
				cmdFlag = append(cmdFlag, &FlagCmd{
					FlagName: fmt.Sprintf("--%s", flg.Name),
					FlagDesc: flg.Usage,
				})
			}
			hlp := &Helper{
				Usage:    "gomigator up seeder [OPTIONS]",
				Argument: cmdFlag,
			}
			printTemplate(helperTmp, hlp)
		}
		upSeeder.Parse(argument[3:])
		command.UpSeeder(c)
	default:
		hlp := &Helper{
			Error: os.Args[2],
		}
		printTemplate(errorTmp, hlp)
	}
}

func createHandle(argument []string, migration, seeder *flag.FlagSet, c *model.Create) {
	switch argument[2] {
	case "migration":
		migration.Usage = func() {
			var cmdFlag []*FlagCmd
			order := []string{"table"}
			for _, item := range order {
				flg := migration.Lookup(item)
				cmdFlag = append(cmdFlag, &FlagCmd{
					FlagName: fmt.Sprintf("--%s", flg.Name),
					FlagDesc: flg.Usage,
				})
			}
			hlp := &Helper{
				Usage:    "gomigator create migration [OPTIONS]",
				Argument: cmdFlag,
			}
			printTemplate(helperTmp, hlp)
		}
		migration.Parse(argument[3:])
		command.CreateMigtaion(c)
	case "seeder":
		seeder.Usage = func() {
			var cmdFlag []*FlagCmd
			order := []string{"table", "name"}
			for _, item := range order {
				flg := seeder.Lookup(item)
				cmdFlag = append(cmdFlag, &FlagCmd{
					FlagName: fmt.Sprintf("--%s", flg.Name),
					FlagDesc: flg.Usage,
				})
			}
			hlp := &Helper{
				Usage:    "gomigator create seeder [OPTIONS]",
				Argument: cmdFlag,
			}
			printTemplate(helperTmp, hlp)
		}
		seeder.Parse(argument[3:])
		command.CreateSeeder(c)
	default:
		hlp := &Helper{
			Error: os.Args[2],
		}
		printTemplate(errorTmp, hlp)
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
