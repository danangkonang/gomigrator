# migration Go-Lang CLI

## On developing 

`gomig` is simple tools database migration for developer
 
- **usage**

```go
  $ go run gomig.go --help
```

- or cli option

```go
  $ go build gomig.go && ./gomig --help
```

- yau can use global usage

```bash
  $ sudo cp gomig /usr/local/bin
```

## example
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
