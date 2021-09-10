# gomigrator

`gomigrator` is a simple tools database migration for developer

## Installation
 
  - local install

  ```bash
  sudo curl -L "https://github.com/danangkonang/gomigrator/releases/download/0.0.9/gomigrator" -o gomigrator && sudo chmod +x gomigrator
  ```

  - global install

  ```bash
  sudo curl -L "https://github.com/danangkonang/gomigrator/releases/download/0.0.9/gomigrator" -o /usr/local/bin/gomigrator
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
* 

## Usage
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

see [example](https://github.com/danangkonang/gomigrator/blob/master/EXAMPLE.md) seeder with fake data
