# migrasion-go-cli

## On developing 

usage

`$ go run gomig.go --help`

or cli option

`$ go build gomig.go && ./gomig --help`

global usage

`$ sudo cp gomig /usr/local/bin`

test

`$ gomig --version`

create migration cli

`$ gomig [option] migration [tablename]`

example

`$ gomig create migration users`

create seeder cli

`$ gomig [option] seeder [tablename]`

example

`$ gomig create seeder users`