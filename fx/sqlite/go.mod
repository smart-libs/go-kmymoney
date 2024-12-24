module github.com/smart-libs/go-kmymoney/fx/sqlite

go 1.22.10

require (
	github.com/smart-libs/go-kmymoney/impl/sqlite v0.0.1
	go.uber.org/fx v1.22.1
)

require (
	github.com/smart-libs/go-kmymoney/spec/lib v0.0.1 // indirect
	go.uber.org/dig v1.17.1 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	go.uber.org/zap v1.26.0 // indirect
	golang.org/x/sys v0.0.0-20220412211240-33da011f77ad // indirect
)

replace (
	github.com/smart-libs/go-kmymoney/impl/sqlite => ../../impl/sqlite
	github.com/smart-libs/go-kmymoney/spec/lib => ../../spec/lib
)
