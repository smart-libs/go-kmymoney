module github.com/smart-libs/go-kmymoney/impl/sqlite

go 1.22.10

require (
	github.com/mattn/go-sqlite3 v1.14.24
	github.com/smart-libs/go-kmymoney/spec/lib v0.0.1
	github.com/stretchr/testify v1.10.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/smart-libs/go-kmymoney/spec/lib => ../../spec/lib
