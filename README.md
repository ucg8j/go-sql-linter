# gsl
> go-sql-linter | good-sql-looks

![](./example/gsl-in-action.gif)

# Dev
```bash
# build
$ go build -o gsl main.go
# test
$ gsl lint example/test.sql
```

# TODOs
**cmd**
- [x] error if filename not `.sql`
- [x] overwrite previous file
- [] take multiple files could be `gsl file1 file2` or `gsl .` with an `-r` flag for recursive

**lint rules**
- [x] Trailing Whitespace
- [x] Greater than one newline
- [x] capitalisation of SQL keywords
  - [] Bug: doesn't work if keyword has delimiter `,` immediately before/after keyword
- [] detecting nested queries
- [] 2 spaces for indentation. Not tabs.
- [] blank line between groups of logic e.g after a select statement, after a from, after Joins etc

**Optimisations**
- [] each lint rule/function loops through lines of the .sql file. Whilst this is O(N), once there are a many of rules e.g. 10 the speed could noticeably be increased
- the in-memory impact of a sql file or several variants is likely very small, however it is an exercise in optimisation using golang features.
