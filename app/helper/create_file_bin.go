package helper

import (
	"fmt"
	"os"
)

func CreateBinFileNew(thisDir, dirMigration string) {
	path := dirMigration + "/bin.go"
	// deteksi apakah file sudah ada
	_, err := os.Stat(path)
	// buat file baru jika belum ada

	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(0)
		}
		writeText := "package main\n\n"
		writeText += "import (\n"
		writeText += "	\"flag\"\n"
		writeText += "	\"fmt\"\n"
		writeText += "	\"io/ioutil\"\n"
		writeText += "	\"os\"\n"
		writeText += "	\"reflect\"\n"
		writeText += "	\"strings\"\n"
		writeText += "	\"text/template\"\n\n"
		writeText += "	\"github.com/danangkonang/" + thisDir + "/db/migration\"\n"
		writeText += "	\"github.com/danangkonang/" + thisDir + "/db/seeder\"\n"
		writeText += "	\"github.com/joho/godotenv\"\n"
		writeText += ")\n\n"

		writeText += "var (\n"
		writeText += "	help    bool\n"
		writeText += "	version bool\n"
		writeText += ")\n\n"

		writeText += "type Tables struct {\n"
		writeText += "	NameTable []string\n"
		writeText += "}\n\n"
		writeText += "var (\n"
		writeText += "	MigrationFolder = \"db/migration\"\n"
		writeText += "	SeederFolder    = \"db/seeder\"\n"
		writeText += `	green           = "\033[32m"`
		writeText += "\n"
		writeText += `	reset           = "\033[0m"`
		writeText += "\n"
		writeText += ")\n\n"

		writeText += "type ComandUsage struct {\n"
		writeText += "	CmdName  string\n"
		writeText += "	CmdAlias string\n"
		writeText += "	CmdDesc  string\n"
		writeText += "}\n\n"

		writeText += "type FlagCmd struct {\n"
		writeText += "	FlagName  string\n"
		writeText += "	FlagAlias string\n"
		writeText += "	FlagDesc  string\n"
		writeText += "}\n\n"

		writeText += "type Helper struct {\n"
		writeText += "	Usage    string\n"
		writeText += "	Version  string\n"
		writeText += "	Error    string\n"
		writeText += "	Option   []*ComandUsage\n"
		writeText += "	Argument []*FlagCmd\n"
		writeText += "}\n\n"

		writeText += "var versionTmp = `Version: {{ .Version }}\n"
		writeText += "`\n\n"

		writeText += "var errorTmp = `unknow comand '{{ .Error }}'\n\n"
		writeText += "see 'gomigator --help'\n"
		writeText += "`\n\n"

		writeText += "var helperTmp = `\n"
		writeText += "Usage: {{ .Usage }}\n"
		writeText += "{{ if .Option }}\n"
		writeText += "Commands:\n"
		writeText += "{{- range .Option}}\n"
		writeText += "	{{ .CmdName }}      {{\"\t\"}}{{ .CmdDesc }}{{ end -}}\n"
		writeText += "{{ end }}\n\n"
		writeText += "Options:\n"
		writeText += "{{- range .Argument}}\n"
		writeText += "	{{ .FlagName }}  {{\"\t\"}}{{ .FlagDesc }}{{ end }}\n\n"
		writeText += "`\n\n"

		writeText += "func printTemplate(temp string, data interface{}) {\n"
		writeText += "	tmpl, err := template.New(\"help\").Parse(temp)\n"
		writeText += "	if err != nil {\n"
		writeText += "		fmt.Println(err.Error())\n"
		writeText += "		os.Exit(0)\n"
		writeText += "	}\n"
		writeText += "	err = tmpl.Execute(os.Stdout, data)\n"
		writeText += "	if err != nil {\n"
		writeText += "		fmt.Println(err.Error())\n"
		writeText += "		os.Exit(0)\n"
		writeText += "	}\n"
		writeText += "	os.Exit(0)\n"
		writeText += "}\n\n"

		writeText += "func globalHelp() {\n"
		writeText += "	hlp := &Helper{\n"
		writeText += "		Usage: \"gomigator [COMAND] [OPTIONS]\",\n"
		writeText += "		Option: []*ComandUsage{\n"
		writeText += "			{\n"
		writeText += "				CmdName: \"init\",\n"
		writeText += "				CmdDesc: \"generate db directory for\",\n"
		writeText += "			},\n"
		writeText += "			{\n"
		writeText += "				CmdName: \"create\",\n"
		writeText += "				CmdDesc: \"create migration or seeder file\",\n"
		writeText += "			},\n"
		writeText += "			{\n"
		writeText += "				CmdName: \"up\",\n"
		writeText += "				CmdDesc: \"exect migration to database\",\n"
		writeText += "			},\n"
		writeText += "			{\n"
		writeText += "				CmdName: \"down\",\n"
		writeText += "				CmdDesc: \"drop migration on databse\",\n"
		writeText += "			},\n"
		writeText += "			{\n"
		writeText += "				CmdName: \"migration\",\n"
		writeText += "				CmdDesc: \"generate type migration\",\n"
		writeText += "			},\n"
		writeText += "			{\n"
		writeText += "				CmdName: \"seeder\",\n"
		writeText += "				CmdDesc: \"generate type seeder\",\n"
		writeText += "			},\n"
		writeText += "		},\n"
		writeText += "		Argument: []*FlagCmd{\n"
		writeText += "			{\n"
		writeText += "				FlagName: \"--table\",\n"
		writeText += "				FlagDesc: \"table name\",\n"
		writeText += "			},\n"
		writeText += "			{\n"
		writeText += "				FlagName: \"--tables\",\n"
		writeText += "				FlagDesc: \"list tables\",\n"
		writeText += "			},\n"
		writeText += "			{\n"
		writeText += "				FlagName: \"--name\",\n"
		writeText += "				FlagDesc: \"generate file name\",\n"
		writeText += "			},\n"
		writeText += "		},\n"
		writeText += "	}\n"
		writeText += "	printTemplate(helperTmp, hlp)\n"
		writeText += "}\n\n"

		writeText += "func init() {\n"
		writeText += "	godotenv.Load()\n"
		writeText += "}\n\n"

		/*
		* Main
		 */
		writeText += "func main() {\n"
		writeText += "	flag.BoolVar(&help, \"h\", false, \"help\")\n"
		writeText += "	flag.BoolVar(&help, \"help\", false, \"help\")\n"
		writeText += "	flag.BoolVar(&version, \"v\", false, \"version\")\n"
		writeText += "	flag.BoolVar(&version, \"version\", false, \"version\")\n"
		writeText += "	flag.Parse()\n"
		writeText += "	if help || len(os.Args[1:]) == 0 {\n"
		writeText += `		globalHelp()`
		writeText += "\n"
		writeText += "	}\n"
		writeText += "	if version {\n"
		writeText += "		hlp := &Helper{\n"
		writeText += "			Version: \"1.1.1\",\n"
		writeText += "		}\n"
		writeText += `		printTemplate(versionTmp, hlp)`
		writeText += "	\n"
		writeText += "	}\n"
		writeText += "	var t Tables\n"
		writeText += "	up := flag.NewFlagSet(\"up\", flag.ExitOnError)\n"
		writeText += "	upMigration := flag.NewFlagSet(\"migration\", flag.ContinueOnError)\n"
		writeText += "	upMigration.Func(\"tables\", \"list file name\", func(s string) error {\n"
		writeText += "		t.NameTable = strings.Fields(s)\n"
		writeText += "		return nil\n"
		writeText += "	})\n"
		writeText += "	upSeeder := flag.NewFlagSet(\"seeder\", flag.ContinueOnError)\n"
		writeText += "	upSeeder.Func(\"tables\", \"list file name\", func(s string) error {\n"
		writeText += "		t.NameTable = strings.Fields(s)\n"
		writeText += "		return nil\n"
		writeText += "	})\n\n"
		writeText += "	down := flag.NewFlagSet(\"down\", flag.ExitOnError)\n"
		writeText += `	downMigration := flag.NewFlagSet("migration", flag.ContinueOnError)`
		writeText += "\n"
		writeText += `	downMigration.Func("tables", "list file name", func(s string) error {`
		writeText += "\n"
		writeText += "		t.NameTable = strings.Fields(s)\n"
		writeText += "		return nil\n"
		writeText += "	})\n"
		writeText += `	downSeeder := flag.NewFlagSet("seeder", flag.ContinueOnError)`
		writeText += "\n"
		writeText += `	downSeeder.Func("tables", "list file name", func(s string) error {`
		writeText += "\n"
		writeText += "		t.NameTable = strings.Fields(s)\n"
		writeText += "		return nil\n"
		writeText += "	})\n"
		writeText += "	if len(os.Args) < 2 {\n"
		writeText += `		globalHelp()`
		writeText += "\n"
		writeText += "	}\n"
		writeText += "	switch os.Args[1] {\n"
		writeText += `	case "up":`
		writeText += "\n"
		writeText += "		up.Parse(os.Args[2:])\n"
		writeText += "		upHandle(os.Args, upMigration, upSeeder, &t)\n"
		writeText += `	case "down":`
		writeText += "\n"
		writeText += "		down.Parse(os.Args[2:])\n"
		writeText += "		downHandle(os.Args, upMigration, upSeeder, &t)\n"
		writeText += "	default:\n"
		writeText += "		hlp := &Helper{\n"
		writeText += "			Error: os.Args[1],\n"
		writeText += "		}\n"
		writeText += "		printTemplate(errorTmp, hlp)\n"
		writeText += "	}\n"
		writeText += "}\n\n"

		writeText += "func downHandle(argument []string, upMigration, upSeeder *flag.FlagSet, tb *Tables) {\n"
		writeText += "	switch argument[2] {\n"
		writeText += `	case "migration":`
		writeText += "\n"
		writeText += "		upMigration.Parse(os.Args[3:])\n"
		writeText += "		migrationDown(tb)\n"
		writeText += `	case "seeder":`
		writeText += "\n"
		writeText += "		upSeeder.Parse(os.Args[3:])\n"
		writeText += "		seederDown(tb)\n"
		writeText += "	default:\n"
		writeText += "		hlp := &Helper{\n"
		writeText += "			Error: os.Args[1],\n"
		writeText += "		}\n"
		writeText += "		printTemplate(errorTmp, hlp)\n"
		writeText += "	}\n"
		writeText += "}\n\n"

		writeText += "func migrationDown(t *Tables) {\n"
		writeText += "	files, err := ioutil.ReadDir(MigrationFolder)\n"
		writeText += "	if err != nil {\n"
		writeText += `		fmt.Println(err.Error())`
		writeText += "\n"
		writeText += "		os.Exit(0)\n"
		writeText += "	}\n"
		writeText += "	if len(t.NameTable) == 0 {\n"
		writeText += "		newFile := []string{}\n"
		writeText += "		for _, file := range files {\n"
		writeText += "			filename := file.Name()\n"
		writeText += `			if filename != "0.go" {`
		writeText += "\n"
		writeText += "				newFile = append(newFile, filename)\n"
		writeText += "			}\n"
		writeText += "		}\n"
		writeText += "		t.NameTable = newFile\n"
		writeText += "	}\n"
		writeText += "	m := migration.Migration{}\n"
		writeText += "	for _, migrate := range t.NameTable {\n"
		// writeText += `		list := strings.Split(migrate, "_migration_")`
		// writeText += "\n"
		// writeText += `		tb_name := strings.Split(list[1], ".go")`
		// writeText += "\n"
		// writeText += `		query := "DROP TABLE IF EXISTS " + tb_name[0] + " CASCADE;"`
		// writeText += "\n"
		// writeText += "		_, err := migration.Connection().Db.Exec(query)\n"
		// writeText += "		if err != nil {\n"
		// writeText += "			fmt.Println(err)\n"
		// writeText += "			os.Exit(0)\n"
		// writeText += "		}\n"
		// writeText += `		fmt.Println(string(green), "success", string(reset), "DROP ", migrate)`
		// writeText += "\n"
		writeText += `		list := strings.Split(migrate, "_migration_")`
		writeText += "\n"
		writeText += `		tb_name := strings.Split(list[1], ".go")`
		writeText += "\n"
		writeText += `		meth := reflect.ValueOf(&m).MethodByName("Down" + strings.Title(tb_name[0]))`
		writeText += "\n"
		writeText += "		meth.Call(nil)\n"
		writeText += "	}\n"
		writeText += "}\n\n"

		writeText += "func seederDown(t *Tables) {\n"
		writeText += "	files, err := ioutil.ReadDir(SeederFolder)\n"
		writeText += "	if err != nil {\n"
		writeText += `		fmt.Println(err.Error())`
		writeText += "\n"
		writeText += "		os.Exit(0)\n"
		writeText += "	}\n"
		writeText += "	if len(t.NameTable) == 0 {\n"
		writeText += "		newFile := []string{}\n"
		writeText += "		for _, file := range files {\n"
		writeText += "			filename := file.Name()\n"
		writeText += `			if filename != "0.go" {`
		writeText += "\n"
		writeText += "				newFile = append(newFile, filename)\n"
		writeText += "			}\n"
		writeText += "		}\n"
		writeText += "		t.NameTable = newFile\n"
		writeText += "	}\n"
		writeText += "	for _, migrate := range t.NameTable {\n"
		writeText += `		list := strings.Split(migrate, "_seeder_")`
		writeText += "\n"
		writeText += `		tb_name := strings.Split(list[1], ".go")`
		writeText += "\n"

		writeText += "		var query string\n"
		writeText += `		if os.Getenv("DB_DRIVER") == "mysql" {`
		writeText += "\n"
		writeText += `			query = "TRUNCATE " + tb_name[0] + " ;"`
		writeText += "\n"
		writeText += "		} else {\n"
		writeText += `			query = "TRUNCATE " + tb_name[0] + " RESTART IDENTITY;"`
		writeText += "\n"
		writeText += "		}\n"
		writeText += "\n"
		writeText += "		_, err := migration.Connection().Db.Exec(query)\n"
		writeText += "		if err != nil {\n"
		writeText += "			fmt.Println(err.Error())\n"
		writeText += "			os.Exit(0)\n"
		writeText += "		}\n"
		writeText += `		fmt.Println(string(green), "success", string(reset), "down ", migrate)`
		writeText += "\n"
		writeText += "	}\n"
		writeText += "}\n\n"

		writeText += "func upHandle(argument []string, upMigration, upSeeder *flag.FlagSet, tb *Tables) {\n"
		writeText += "	switch argument[2] {\n"
		writeText += `	case "migration":`
		writeText += "\n"
		writeText += "		upMigration.Parse(os.Args[3:])\n"
		writeText += "		migrationUp(tb)\n"
		writeText += `	case "seeder":`
		writeText += "\n"
		writeText += "		upSeeder.Parse(os.Args[3:])\n"
		writeText += "		seederUp(tb)\n"
		writeText += "	default:\n"
		writeText += "		hlp := &Helper{\n"
		writeText += "			Error: os.Args[1],\n"
		writeText += "		}\n"
		writeText += "		printTemplate(errorTmp, hlp)\n"
		writeText += "	}\n"
		writeText += "}\n\n"

		writeText += "func migrationUp(t *Tables) {\n"
		writeText += "	files, err := ioutil.ReadDir(MigrationFolder)\n"
		writeText += "	if err != nil {\n"
		writeText += `		fmt.Println(err.Error())`
		writeText += "\n"
		writeText += "		os.Exit(0)\n"
		writeText += "	}\n"
		writeText += "	if len(t.NameTable) == 0 {\n"
		writeText += "		newFile := []string{}\n"
		writeText += "		for _, file := range files {\n"
		writeText += "			filename := file.Name()\n"
		writeText += `			if filename != "0.go" {`
		writeText += "\n"
		writeText += "				newFile = append(newFile, filename)\n"
		writeText += "			}\n"
		writeText += "		}\n"
		writeText += "		t.NameTable = newFile\n"
		writeText += "	}\n"
		writeText += "	m := migration.Migration{}\n"
		writeText += "	for _, migrate := range t.NameTable {\n"
		writeText += `		list := strings.Split(migrate, "_migration_")`
		writeText += "\n"
		writeText += `		tb_name := strings.Split(list[1], ".go")`
		writeText += "\n"
		writeText += `		meth := reflect.ValueOf(&m).MethodByName("Up" + strings.Title(tb_name[0]))`
		writeText += "\n"
		writeText += "		meth.Call(nil)\n"
		writeText += "	}\n"
		writeText += "}\n\n"

		writeText += "func seederUp(t *Tables) {\n"
		writeText += "	files, err := ioutil.ReadDir(SeederFolder)\n"
		writeText += "	if err != nil {\n"
		writeText += `		fmt.Println(err.Error())`
		writeText += "\n"
		writeText += "		os.Exit(0)\n"
		writeText += "	}\n"
		writeText += "	if len(t.NameTable) == 0 {\n"
		writeText += "		newFile := []string{}\n"
		writeText += "		for _, file := range files {\n"
		writeText += "			filename := file.Name()\n"
		writeText += `			if filename != "0.go" {`
		writeText += "\n"
		writeText += "				newFile = append(newFile, filename)\n"
		writeText += "			}\n"
		writeText += "		}\n"
		writeText += "		t.NameTable = newFile\n"
		writeText += "	}\n"
		writeText += "	s := seeder.Seeder{}\n"
		writeText += "	for _, seed := range t.NameTable {\n"
		writeText += `		list := strings.Split(seed, "_seeder_")`
		writeText += "\n"
		writeText += `		func_name := strings.Split(list[1], ".go")`
		writeText += "\n"
		writeText += "		meth := reflect.ValueOf(&s).MethodByName(strings.Title(func_name[0]))\n"
		writeText += "		meth.Call(nil)\n"
		writeText += "	}\n"
		writeText += "}\n\n"

		file.WriteString(writeText)
		defer file.Close()
	}
}
