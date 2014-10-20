
# go-pg-escape

 Escape Postgres queries.

 View the [docs](http://godoc.org/github.com/tj/go-pg-escape).

## Installation

```
$ go get github.com/tj/go-pg-escape
```

## Example

```go
s := Escape("SELECT %I FROM %I WHERE %I=%L", "some stuff", "some table", "some column", "some value")
exp := `SELECT "some stuff" FROM "some table" WHERE "some column"='some value'`
assert.Equal(t, exp, s)
```

# License

 MIT