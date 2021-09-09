# gomigrator

`gomigrator` is a simple tools database migration for developer

## Installation
 
  - local install

  ```bash
  sudo curl -L "https://github.com/danangkonang/gomigrator/releases/download/0.0.8/gomigrator" -o gomigrator && chmod +x gomigrator
  ```

  - global install

  ```bash
  sudo curl -L "https://github.com/danangkonang/gomigrator/releases/download/0.0.8/gomigrator" -o /usr/local/bin/gomigrator
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
- ***Migration***

```bash
  # create migration
  gomigrator create migration --table [table name]

```

- ***Seeder***

```bash
  # create seeder file
  gomigrator create seeder --table [table name]

  # runing seeder
  gomigrator up seeder
```

- **Up**

```bash
  # up all migration
  gomigrator up migration

```

- **Down**

```bash
  # down 
  gomigrator down migration

  # or spesifik table 
  gomigrator down migration --tables [list migration file]

````
