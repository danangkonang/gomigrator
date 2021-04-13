# migrasion-go-cli

## On developing 

usage

```bash
  $ go run gomig.go --help
```

or cli option

```bash
  $ go build gomig.go && ./gomig --help
```

global usage

```bash
  $ sudo cp gomig /usr/local/bin
```

## example
create migration

```bash
  $ gomig create migration -t [table name]
```

create seeder cli

```bash
  $ gomig create seeder -n [seeder name] -t [table name]
```


runing migration

```bash
  $ gomig run migration
```

runing seeder

```bash
  $ gomig run seeder
```

undo migration

```bash
  $ gomig drop
````

undo seeder

```bash
  $ gomig down
````