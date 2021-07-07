# migration Go-Lang CLI

## On developing 

`gomig` is simple tools database migration for developer

## Installation
 
  - Run this command to download:

  ```bash
    $ sudo curl -L "https://github.com/danangkonang/migration-go-cli/releases/download/v0.0.5/gomig" -o /usr/local/bin/gomig
  ```

  - Apply executable permissions to the binary:
  ```bash
    $ sudo chmod +x /usr/local/bin/gomig
  ```

  -  If the command gomig fails you cant run this
  ```bash
    $ sudo ln -s /usr/local/bin/gomig /usr/bin/gomig
  ```

  - Test installation.
  ```bash
    $ goming --version
  ```

## Databases

* PostgreSQL
* 

## Example
- ***migration***

```bash
  # create migration
  $ gomig create migration -t [table name]

  # running all migration
  $ gomig run migration
```

- ***seeder***

```bash
  # create seeder file
  $ gomig create seeder -n [seeder name] -t [table name]

  # runing seeder
  $ gomig run seeder
```

- **undo**

```bash
  # delete all tables 
  $ gomig drop

  # or spesifik table 
  $ gomig drop [table1 table2 ..]

  # delete all data seeder
  $ gomig reset

  # or spesifik table 
  $ gomig reset [table1 table2 ..]
````
