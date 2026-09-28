# dup-schema-repro

Minimal reproduction for a [pb33f/libopenapi](https://github.com/pb33f/libopenapi) bundler issue. Details are in the [linked upstream issue](https://github.com/pb33f/libopenapi/issues/644).

## Run

Requires Go 1.26.7. The `go` command downloads it if your install is older.

```bash
go run .
```

To run against a different libopenapi version:

```bash
go get github.com/pb33f/libopenapi@v0.36.0
go run .
```