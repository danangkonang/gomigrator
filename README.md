# migration Go-Lang CLI

## On developing 

`gomigrator` is a simple tools database migration for developer

## Installation
 
  - Run this command to download:

  ```bash
  #local isntall
  sudo curl -L "https://github.com/danangkonang/migration-go-cli/releases/download/0.0.8/main" -o gomigrator && chmod +x gomigrator

  #global install
  sudo curl -L "https://github.com/danangkonang/migration-go-cli/releases/download/0.0.8/main" -o /usr/local/bin/gomigrator
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

* PostgreSQL
* Mysql

## Basic
- ***migration***

```bash
  # create migration
  gomigrator create migration --table [table name]

```

- ***seeder***

```bash
  # create seeder file
  gomigrator create seeder --table [table name]

  # runing seeder
  gomigrator up seeder
```

- **up**

```bash
  # up all migration
  gomigrator up migration

```

- **down**

```bash
  # down 
  gomigrator down migration

  # or spesifik table 
  gomigrator down migration --tables [list migration file]

````
