package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/danangkonang/gomigrator/app/command"
	"github.com/danangkonang/gomigrator/app/model"
	"github.com/danangkonang/gomigrator/templates"
	"github.com/joho/godotenv"
)

var (
	help    bool
	version bool
)

func init() {
	godotenv.Load()
}

func globalHelp() {
	hlp := &templates.Helper{
		Usage: "gomigator [COMAND] [OPTIONS]",
		Option: []*templates.ComandUsage{
			{
				CmdName: "init",
				CmdDesc: "Generate default directory to manage migration",
			},
			{
				CmdName: "create",
				CmdDesc: "Creates new migration or seeder file",
			},
			{
				CmdName: "up",
				CmdDesc: "Migrate to database",
			},
			{
				CmdName: "down",
				CmdDesc: "Roll back migration",
			},
			{
				CmdName: "migration",
				CmdDesc: "Generate migration type",
			},
			{
				CmdName: "seeder",
				CmdDesc: "Generate seeder type",
			},
		},
		Argument: []*templates.FlagCmd{
			{
				FlagName: "-h, --help",
				FlagDesc: "Print help gomigrator",
			},
			{
				FlagName: "-v, --version",
				FlagDesc: "Print version gomigrator",
			},
			{
				FlagName: "--table",
				FlagDesc: "Table name",
			},
			{
				FlagName: "--tables",
				FlagDesc: "List table",
			},
			{
				FlagName: "--name",
				FlagDesc: "Generate file name",
			},
		},
	}
	templates.PrintTemplate(templates.HelperTmp, hlp)
}

func run() {
	flag.BoolVar(&help, "h", false, "help")
	flag.BoolVar(&help, "help", false, "help")
	flag.BoolVar(&version, "v", false, "version")
	flag.BoolVar(&version, "version", false, "version")
	flag.Parse()
	if help || len(os.Args[1:]) == 0 {
		globalHelp()
	}
	if version {
		hlp := &templates.Helper{
			Version: "0.1.2",
		}
		// printTemplate(versionTmp, hlp)
		templates.PrintTemplate(templates.VersionTmp, hlp)
	}
	var i model.Init

	/*
	*
	 */
	init := flag.NewFlagSet("init", flag.ExitOnError)
	init.Func("driver", "Set up db driver", func(s string) error {
		i.Driver = s
		return nil
	})
	init.Func("port", "Set up db port", func(s string) error {
		port, _ := strconv.Atoi(s)
		i.Port = port
		return nil
	})
	init.Func("host", "Set up db host", func(s string) error {
		i.Host = s
		return nil
	})
	init.Func("database", "Set up db host", func(s string) error {
		i.DAtabase = s
		return nil
	})
	init.Func("user", "Set up db user", func(s string) error {
		i.User = s
		return nil
	})
	init.Func("password", "Set up db password", func(s string) error {
		i.Password = s
		return nil
	})

	/*
	*
	 */
	create := flag.NewFlagSet("create", flag.ExitOnError)
	var c model.Create
	migration := flag.NewFlagSet("migration", flag.ContinueOnError)
	migration.Func("table", "Generate file name", func(s string) error {
		c.TableName = s
		return nil
	})
	seeder := flag.NewFlagSet("seeder", flag.ContinueOnError)
	seeder.Func("name", "Generate file name", func(s string) error {
		c.FileName = s
		return nil
	})
	seeder.Func("table", "Generate file name", func(s string) error {
		c.TableName = s
		return nil
	})

	/*
	*
	 */
	up := flag.NewFlagSet("up", flag.ExitOnError)
	var upM model.UpDown
	upMigration := flag.NewFlagSet("migration", flag.ContinueOnError)
	upMigration.Func("tables", "List migration file name", func(s string) error {
		upM.Tables = strings.Fields(s)
		return nil
	})
	upSeeder := flag.NewFlagSet("seeder", flag.ContinueOnError)
	upSeeder.Func("tables", "List seeder file name", func(s string) error {
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
	downMigration.Func("tables", "List migration file name", func(s string) error {
		dwM.Tables = strings.Fields(s)
		return nil
	})

	// remove seeder data
	downSeeder := flag.NewFlagSet("seeder", flag.ContinueOnError)
	downSeeder.Func("tables", "List seeder file name", func(s string) error {
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
			var cmdFlag []*templates.FlagCmd
			for _, item := range order {
				flg := init.Lookup(item)
				cmdFlag = append(cmdFlag, &templates.FlagCmd{
					FlagName: fmt.Sprintf("--%s", flg.Name),
					FlagDesc: flg.Usage,
				})
			}
			hlp := &templates.Helper{
				Usage:    "gomigator init [OPTIONS]",
				Argument: cmdFlag,
			}
			templates.PrintTemplate(templates.HelperTmp, hlp)
		}
		init.Parse(os.Args[2:])
		command.Init(&i)
	case "create":
		create.Usage = func() {
			var cmdOption []*templates.ComandUsage
			cmdOption = append(cmdOption, &templates.ComandUsage{
				CmdName: migration.Name(),
				CmdDesc: "migration type",
			})
			cmdOption = append(cmdOption, &templates.ComandUsage{
				CmdName: seeder.Name(),
				CmdDesc: "migration type",
			})
			var cmdFlag []*templates.FlagCmd
			order := []string{"table"}
			for _, item := range order {
				flg := migration.Lookup(item)
				cmdFlag = append(cmdFlag, &templates.FlagCmd{
					FlagName: fmt.Sprintf("--%s", flg.Name),
					FlagDesc: flg.Usage,
				})
			}
			hlp := &templates.Helper{
				Usage:    "gomigator create [COMMAND] [OPTIONS]",
				Option:   cmdOption,
				Argument: cmdFlag,
			}
			templates.PrintTemplate(templates.HelperTmp, hlp)
		}
		if len(os.Args) <= 2 {
			globalHelp()
		}
		create.Parse(os.Args[2:])
		createHandle(os.Args, migration, seeder, &c)
	case "up":
		up.Usage = func() {
			var cmdOption []*templates.ComandUsage
			cmdOption = append(cmdOption, &templates.ComandUsage{
				CmdName: upMigration.Name(),
				CmdDesc: "migration type",
			})
			cmdOption = append(cmdOption, &templates.ComandUsage{
				CmdName: upSeeder.Name(),
				CmdDesc: "migration type",
			})
			var cmdFlag []*templates.FlagCmd
			order := []string{"tables"}
			for _, item := range order {
				flg := upMigration.Lookup(item)
				cmdFlag = append(cmdFlag, &templates.FlagCmd{
					FlagName: fmt.Sprintf("--%s", flg.Name),
					FlagDesc: flg.Usage,
				})
			}
			hlp := &templates.Helper{
				Usage:    "gomigator up [COMMAND] [OPTIONS]",
				Option:   cmdOption,
				Argument: cmdFlag,
			}
			templates.PrintTemplate(templates.HelperTmp, hlp)
		}
		if len(os.Args) <= 2 {
			globalHelp()
		}
		up.Parse(os.Args[2:])
		upHandle(os.Args, upMigration, upSeeder, &upM)
	case "down":
		down.Usage = func() {
			var cmdOption []*templates.ComandUsage
			cmdOption = append(cmdOption, &templates.ComandUsage{
				CmdName: upMigration.Name(),
				CmdDesc: "migration type",
			})
			cmdOption = append(cmdOption, &templates.ComandUsage{
				CmdName: upSeeder.Name(),
				CmdDesc: "migration type",
			})
			var cmdFlag []*templates.FlagCmd
			order := []string{"tables"}
			for _, item := range order {
				flg := upMigration.Lookup(item)
				cmdFlag = append(cmdFlag, &templates.FlagCmd{
					FlagName: fmt.Sprintf("--%s", flg.Name),
					FlagDesc: flg.Usage,
				})
			}
			hlp := &templates.Helper{
				Usage:    "gomigator down [COMMAND] [OPTIONS]",
				Option:   cmdOption,
				Argument: cmdFlag,
			}
			templates.PrintTemplate(templates.HelperTmp, hlp)
		}
		if len(os.Args) <= 2 {
			globalHelp()
		}
		down.Parse(os.Args[2:])
		downHandle(os.Args, downMigration, downSeeder, &dwM)
	default:
		hlp := &templates.Helper{
			Error: fmt.Sprintf("unknow command: %s", os.Args[1]),
		}
		templates.PrintTemplate(templates.ErrorTmp, hlp)
	}
}

func downHandle(argument []string, downMigration, downSeeder *flag.FlagSet, c *model.UpDown) {
	switch argument[2] {
	case "migration":
		downMigration.Usage = func() {
			var cmdFlag []*templates.FlagCmd
			order := []string{"tables"}
			for _, item := range order {
				flg := downMigration.Lookup(item)
				cmdFlag = append(cmdFlag, &templates.FlagCmd{
					FlagName: fmt.Sprintf("--%s", flg.Name),
					FlagDesc: flg.Usage,
				})
			}
			hlp := &templates.Helper{
				Usage:    "gomigator down migration [OPTIONS]",
				Argument: cmdFlag,
			}
			templates.PrintTemplate(templates.HelperTmp, hlp)
		}
		downMigration.Parse(argument[3:])
		command.DownMigration(c)
	case "seeder":
		downSeeder.Usage = func() {
			var cmdFlag []*templates.FlagCmd
			order := []string{"tables"}
			for _, item := range order {
				flg := downSeeder.Lookup(item)
				cmdFlag = append(cmdFlag, &templates.FlagCmd{
					FlagName: fmt.Sprintf("--%s", flg.Name),
					FlagDesc: flg.Usage,
				})
			}
			hlp := &templates.Helper{
				Usage:    "gomigator down seeder [OPTIONS]",
				Argument: cmdFlag,
			}
			templates.PrintTemplate(templates.HelperTmp, hlp)
		}
		downSeeder.Parse(argument[3:])
		command.DownSeeder(c)
	default:
		hlp := &templates.Helper{
			Error: fmt.Sprintf("unknow command: %s", os.Args[2]),
		}
		templates.PrintTemplate(templates.ErrorTmp, hlp)
	}
}

func upHandle(argument []string, upMigration, upSeeder *flag.FlagSet, c *model.UpDown) {
	switch argument[2] {
	case "migration":
		upMigration.Usage = func() {
			var cmdFlag []*templates.FlagCmd
			order := []string{"tables"}
			for _, item := range order {
				flg := upMigration.Lookup(item)
				cmdFlag = append(cmdFlag, &templates.FlagCmd{
					FlagName: fmt.Sprintf("--%s", flg.Name),
					FlagDesc: flg.Usage,
				})
			}
			hlp := &templates.Helper{
				Usage:    "gomigator up migration [OPTIONS]",
				Argument: cmdFlag,
			}
			templates.PrintTemplate(templates.HelperTmp, hlp)
		}
		upMigration.Parse(argument[3:])
		command.UpMigration(c)
	case "seeder":
		upSeeder.Usage = func() {
			var cmdFlag []*templates.FlagCmd
			order := []string{"tables"}
			for _, item := range order {
				flg := upSeeder.Lookup(item)
				cmdFlag = append(cmdFlag, &templates.FlagCmd{
					FlagName: fmt.Sprintf("--%s", flg.Name),
					FlagDesc: flg.Usage,
				})
			}
			hlp := &templates.Helper{
				Usage:    "gomigator up seeder [OPTIONS]",
				Argument: cmdFlag,
			}
			templates.PrintTemplate(templates.HelperTmp, hlp)
		}
		upSeeder.Parse(argument[3:])
		command.UpSeeder(c)
	default:
		hlp := &templates.Helper{
			Error: os.Args[2],
		}
		templates.PrintTemplate(templates.HelperTmp, hlp)
	}
}

func createHandle(argument []string, migration, seeder *flag.FlagSet, c *model.Create) {
	switch argument[2] {
	case "migration":
		migration.Usage = func() {
			var cmdFlag []*templates.FlagCmd
			order := []string{"table"}
			for _, item := range order {
				flg := migration.Lookup(item)
				cmdFlag = append(cmdFlag, &templates.FlagCmd{
					FlagName: fmt.Sprintf("--%s", flg.Name),
					FlagDesc: flg.Usage,
				})
			}
			hlp := &templates.Helper{
				Usage:    "gomigator create migration [OPTIONS]",
				Argument: cmdFlag,
			}
			templates.PrintTemplate(templates.HelperTmp, hlp)
		}
		migration.Parse(argument[3:])
		command.CreateMigtaion(c)
	case "seeder":
		seeder.Usage = func() {
			var cmdFlag []*templates.FlagCmd
			order := []string{"table", "name"}
			for _, item := range order {
				flg := seeder.Lookup(item)
				cmdFlag = append(cmdFlag, &templates.FlagCmd{
					FlagName: fmt.Sprintf("--%s", flg.Name),
					FlagDesc: flg.Usage,
				})
			}
			hlp := &templates.Helper{
				Usage:    "gomigator create seeder [OPTIONS]",
				Argument: cmdFlag,
			}
			templates.PrintTemplate(templates.HelperTmp, hlp)
		}
		seeder.Parse(argument[3:])
		command.CreateSeeder(c)
	default:
		hlp := &templates.Helper{
			Error: os.Args[2],
		}
		templates.PrintTemplate(templates.ErrorTmp, hlp)
	}
}

func main() {
	run()
	// if len("newFile") > 0 {
	// 	// app.FileName = app.FileName + fmt.Sprintf("%03d", len(newFile))
	// 	// app.FileName = app.FileName + fmt.Sprintf("%d", time.Now().Unix())
	// 	fmt.Println(fmt.Sprintf("%03d", len("ini")))
	// }
}
