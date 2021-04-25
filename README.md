# gsl
> go-sql-linter | good-sql-looks

# Dev
```bash
# build
$ go build -o gsl main.go
# test
$ gsl lint example/test.sql
```

# TODOs
**cmd**
- [] error if filename not `.sql`
- [] overwrite previous file

**lint rules**

- [x] Trailing Whitespace
- [] Greater than one newline
- [] capitalisation of SQL keywords
...
