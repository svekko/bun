module github.com/uptrace/bun/example/belongs-to

go 1.16

replace github.com/uptrace/bun => ../..

replace github.com/uptrace/bun/extra/bundebug => ../../extra/bundebug

replace github.com/uptrace/bun/dialect/sqlitedialect => ../../dialect/sqlitedialect

replace github.com/uptrace/bun/driver/sqliteshim => ../../driver/sqliteshim

require (
	github.com/mattn/go-colorable v0.1.11 // indirect
	github.com/uptrace/bun v1.0.11
	github.com/uptrace/bun/dialect/sqlitedialect v1.0.11
	github.com/uptrace/bun/driver/sqliteshim v1.0.11
	github.com/uptrace/bun/extra/bundebug v1.0.11
)
