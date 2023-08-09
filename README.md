# sqlpassctxcheck

sqlpassctxcheck is a program for checking for sql module method call without ctx.
Using this tool, you can avoid falling outside of distributed tracing by forgetting to pass the context.

## Install

```
go install github.com/400f/sqlpassctxcheck/cmd/sqlpassctxcheck@latest
```

## Usage

```go
package main

import "database/sql"

func Call(db *sql.DB)  {
  rows, err := db.Query("SELECT * FROM foo")
  if err != nil {
    return err
  }

  // ...
}
```

```
$ go vet -vettool=(which sqlpassctxcheck) ./...

main.go:6:15 use (*database/sql.DB).QueryContext instead of (*database/sql.DB).Query
```

## Develop
```sh
$ go test -v ./
```
