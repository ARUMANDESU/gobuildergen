# gobuildergen

Generates builder pattern code for Go structs. That's it.

## Install

```
go install github.com/ARUMANDESU/gobuildergen@latest
```

## Usage

Add `//go:generate` above your struct:

```go
//go:generate gobuildergen --type MyStruct
type MyStruct struct {
    Name    string
    Age     int
    Email   string `builder:"default=\"unknown\""`
    secret  string `builder:"-"`
}
```

Run:

```
go generate ./...
```

This creates `myfile_builder_gen.go` with a fluent builder:

```go
s := NewMyStructBuilder().
    WithDefault().
    Name("John").
    Age(30).
    Build()
```

## Struct tags

- `builder:"default=value"` — sets a default value (applied via `WithDefault()`)
- `builder:"-"` — skips the field entirely
