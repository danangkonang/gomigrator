# gomigrator

`gomigrator` is a simple tools database migration for developer

## Installation
 
  - local install

  ```bash
  sudo curl -L "https://github.com/danangkonang/gomigrator/releases/download/0.1.2/gomigrator" -o gomigrator && sudo chmod +x gomigrator
  ```

  - global install

  ```bash
  sudo curl -L "https://github.com/danangkonang/gomigrator/releases/download/0.1.2/gomigrator" -o /usr/local/bin/gomigrator
  ```

  - Apply executable permissions to the binary:
  ```bash
  sudo chmod +x /usr/local/bin/gomigrator
  ```

  -  If the command gomigrator fails you cant run this
  ```bash
  sudo ln -s /usr/local/bin/gomigrator /usr/bin/gomigrator
  ```

  - Test installation.
  ```bash
  gomigrator --version
  ```

## Databases

* PSQL
* MySQL
* 

## Usage
- ***Help**

```bash
Usage: gomigator [COMAND] [OPTIONS]

Commands:
  init          Generate default directory to manage migration
  create        Creates new migration or seeder file
  up            Migrate to database
  down          Roll back migration
  migration     Generate migration type
  seeder        Generate seeder type

Options:
  -h, --help            Print help gomigrator
  -v, --version         Print version gomigrator
  --table               Table name
  --tables              List table
  --name                Generate file name
```
- ***Init***

```bash
  gomigrator init
```
Default init will generate driver postgres, you can custom driver with `gomigrator init --driver [mysql/psql]` or other flags  `gomigrator init --help`.

gomigrator saved connection file `db/mogration/0.go`

- ***Migration***

```bash
  # create migration
  gomigrator create migration --table [table name]

  # up migration
  gomigrator up migration

  # down migration
  gomigrator down migration
```

- ***Seeder***

```bash
  # create seeder file
  gomigrator create seeder --table [table name]

  # up seeder
  gomigrator up seeder

  # down seeder
  gomigrator down seeder
```

See [example](https://github.com/danangkonang/gomigrator/blob/master/EXAMPLE.md) seeder with fake data
